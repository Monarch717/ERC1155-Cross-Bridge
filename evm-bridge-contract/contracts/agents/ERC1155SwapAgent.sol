// SPDX-License-Identifier: MIT
pragma solidity ^0.8.2;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC1155/utils/ERC1155HolderUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/introspection/ERC165CheckerUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "../tokens/MirroredERC1155.sol";
import "../interfaces/IERC1155Mirrored.sol";

contract ERC1155SwapAgent is
    Initializable,
    OwnableUpgradeable,
    ERC1155HolderUpgradeable
{
    using ERC165CheckerUpgradeable for address;

    // -- ERC1155 INTERFACE ID --
    bytes4 private constant _ERC_1155_INTERFACE_ID = 0xd9b67a26;

    // -- Error Constants --
    // ERC1155SwapAgent::registerSwapPair:: token does not conform ERC1155 standard 
    string private constant _ERR1155_REGISTER_NOT_1155 = "A1155.1155.1.1";
    // ERC1155SwapAgent::registerSwapPair:: token is already registered
    string private constant _ERR1155_REGISTER_TOKEN_EXISTS = "A1155.1155.1.2";
    // ERC1155SwapAgent::registerSwapPair:: empty token name
    string private constant _ERR1155_REGISTER_EMPTY_TOKEN_NAME = "A1155.1155.1.3";
    // ERC1155SwapAgent::registerSwapPair:: empty token symbol
    string private constant _ERR1155_REGISTER_EMPTY_TOKEN_SYMBOL = "A1155.1155.1.4";

    // ERC1155SwapAgent::createSwapPair:: mirrored token is already deployed
    string private constant _ERR1155_CREATE_MIRRORED_EXISTS = "A1155.1155.2.1";

    // ERC1155SwapAgent::swap:: token has no swap pair
    string private constant _ERR1155_SWAP_NO_PAIR = "A1155.1155.3.1";

    // ERC1155SwapAgent::fill:: tx hash was already filled
    string private constant _ERR1155_FILL_ALREADY_FILLED = "A1155.1155.4.1";
    // ERC1155SwapAgent::fill:: token has no swap pair
    string private constant _ERR1155_FILL_NO_PAIR = "A1155.1155.4.2";

    // -- Storage Variables --
    mapping(uint256 => mapping(address => bool)) public registeredToken;
    mapping(bytes32 => bool) public filledSwap;
    mapping(uint256 => mapping(address => address)) public swapMappingIncoming;
    mapping(uint256 => mapping(address => address)) public swapMappingOutgoing;

    /* Events from outgoing messages */
    event SwapPairRegister(
        address indexed sponsor,
        address indexed tokenAddress,
        uint256 toChainId,
        uint256 feeAmount
    );

    event SwapStarted(
        address indexed tokenAddr,
        address indexed sender,
        address indexed recipient,
        uint256 dstChainId,
        uint256[] ids,
        uint256[] amounts,
        uint256 feeAmount
    );

    event BackwardSwapStarted(
        address indexed mirroredTokenAddr,
        address indexed sender,
        address indexed recipient,
        uint256 dstChainId,
        uint256[] ids,
        uint256[] amounts,
        uint256 feeAmount
    );

    /* Events from incoming messages */
    event SwapPairCreated(
        bytes32 indexed registerTxHash,
        address indexed fromTokenAddr,
        address indexed mirroredTokenAddr,
        uint256 fromChainId
    );

    event SwapFilled(
        bytes32 indexed swapTxHash,
        address indexed fromTokenAddr,
        address indexed recipient,
        address mirroredTokenAddr,
        uint256 fromChainId,
        uint256[] ids,
        uint256[] amounts
    );

    event BackwardSwapFilled(
        bytes32 indexed swapTxHash,
        address indexed tokenAddr,
        address indexed recipient,
        uint256 fromChainId,
        uint256[] ids,
        uint256[] amounts
    );

    function initialize() public initializer {
        __Ownable_init();
    }

    function createSwapPair(
        bytes32 registerTxHash,
        address fromTokenAddr,
        uint256 fromChainId,
        string calldata uri
    ) public onlyOwner {
        require(
            swapMappingIncoming[fromChainId][fromTokenAddr] == address(0x0),
            _ERR1155_CREATE_MIRRORED_EXISTS
        );

        MirroredERC1155 mirrored = new MirroredERC1155();
        mirrored.initialize(uri);

        swapMappingIncoming[fromChainId][fromTokenAddr] = address(mirrored);
        swapMappingOutgoing[fromChainId][address(mirrored)] = fromTokenAddr;

        emit SwapPairCreated(
            registerTxHash,
            fromTokenAddr,
            address(mirrored),
            fromChainId
        );
    }

    function registerSwapPair(address tokenAddr, uint256 chainId)
        external
        payable
    {
        require(
          tokenAddr.supportsInterface(_ERC_1155_INTERFACE_ID),
          _ERR1155_REGISTER_NOT_1155
        );

        require(
            !registeredToken[chainId][tokenAddr],
            _ERR1155_REGISTER_TOKEN_EXISTS
        );

        registeredToken[chainId][tokenAddr] = true;

        if (msg.value != 0) {
            payable(owner()).transfer(msg.value);
        }

        emit SwapPairRegister(
            msg.sender,
            tokenAddr,
            chainId,
            msg.value
        );
    }

    function swap(
        address tokenAddr,
        address recipient,
        uint256[] calldata ids,
        uint256[] calldata amounts,
        uint256 dstChainId
    ) external payable {
        if (msg.value != 0) {
            payable(owner()).transfer(msg.value);
        }

        // try forward swap
        if (registeredToken[dstChainId][tokenAddr]) {
            IERC1155 token = IERC1155(tokenAddr);
            token.safeBatchTransferFrom(msg.sender, address(this), ids, amounts, "0x00");

            emit SwapStarted(
                tokenAddr,
                msg.sender,
                recipient,
                dstChainId,
                ids,
                amounts,
                msg.value
            );

            return;
        }

        // try backward swap
        address dstTokenAddr = swapMappingOutgoing[dstChainId][tokenAddr];
        if (dstTokenAddr != address(0x0)) {
            IERC1155Mirrored mirroredToken = IERC1155Mirrored(tokenAddr);
            mirroredToken.safeBatchTransferFrom(msg.sender, address(this), ids, amounts, "0x00");
            mirroredToken.burnBatch(address(this), ids, amounts);

            emit BackwardSwapStarted(
                tokenAddr,
                msg.sender,
                recipient,
                dstChainId,
                ids,
                amounts,
                msg.value
            );

            return;
        }

        revert(_ERR1155_SWAP_NO_PAIR);
    }

    function fill(
        bytes32 swapTxHash,
        address fromTokenAddr,
        address recipient,
        uint256 fromChainId,
        uint256[] calldata ids,
        uint256[] calldata amounts
    ) public onlyOwner {
        require(!filledSwap[swapTxHash], _ERR1155_FILL_ALREADY_FILLED);
        filledSwap[swapTxHash] = true;

        // fill forward swap, it means our core server will find the related mirrored token
        // and assign the value to fromTokenAddr
        address mirroredTokenAddr = swapMappingIncoming[fromChainId][
            fromTokenAddr
        ];
        if (mirroredTokenAddr != address(0x0)) {
            IERC1155Mirrored mirroredToken = IERC1155Mirrored(mirroredTokenAddr);
            mirroredToken.mintBatch(recipient, ids, amounts, "0x00");

            emit SwapFilled(
                swapTxHash,
                fromTokenAddr,
                recipient,
                mirroredTokenAddr,
                fromChainId,
                ids,
                amounts
            );

            return;
        }

        // fill backward swap, it means that this token is the one users have been sent before
        // our server will find this token from the given mirrored token in the BackwardSwapStarted event
        // and assign the value to fromTokenAddr
        if (registeredToken[fromChainId][fromTokenAddr]) {
            IERC1155 token = IERC1155(fromTokenAddr);
            token.safeBatchTransferFrom(address(this), recipient, ids, amounts, "0x00");

            emit BackwardSwapFilled(
                swapTxHash,
                fromTokenAddr,
                recipient,
                fromChainId,
                ids,
                amounts
            );

            return;
        }

        revert(_ERR1155_FILL_NO_PAIR);
    }
}
