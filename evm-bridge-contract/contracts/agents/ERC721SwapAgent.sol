// SPDX-License-Identifier: MIT
pragma solidity ^0.8.2;

import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC721/utils/ERC721HolderUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/utils/introspection/ERC165CheckerUpgradeable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "../tokens/MirroredERC721.sol";
import "../interfaces/IERC721Mirrored.sol";

contract ERC721SwapAgent is
    Initializable,
    OwnableUpgradeable,
    ERC721HolderUpgradeable
{
    using ERC165CheckerUpgradeable for address;

    // -- ERC721 INTERFACE ID --
    bytes4 private constant _ERC_721_INTERFACE_ID = 0x80ac58cd;

    // -- Error Constants --
    // ERC721SwapAgent::registerSwapPair:: token does not conform ERC721 standard
    string private constant _ERR721_REGISTER_NOT_721 = "A721.721.1.1";
    // ERC721SwapAgent::registerSwapPair:: token is already registered
    string private constant _ERR721_REGISTER_TOKEN_EXISTS = "A721.721.1.2";
    // ERC721SwapAgent::registerSwapPair:: empty token name
    string private constant _ERR721_REGISTER_EMPTY_TOKEN_NAME = "A721.721.1.3";
    // ERC721SwapAgent::registerSwapPair:: empty token symbol
    string private constant _ERR721_REGISTER_EMPTY_TOKEN_SYMBOL = "A721.721.1.4";

    // ERC721SwapAgent::createSwapPair:: mirrored token is already deployed
    string private constant _ERR721_CREATE_MIRRORED_EXISTS = "A721.721.2.1";

    // ERC721SwapAgent::swap:: token has no swap pair
    string private constant _ERR721_SWAP_NO_PAIR = "A721.721.3.1";

    // ERC721SwapAgent::fill:: tx hash was already filled
    string private constant _ERR721_FILL_ALREADY_FILLED = "A721.721.4.1";
    // ERC721SwapAgent::fill:: token has no swap pair
    string private constant _ERR721_FILL_NO_PAIR = "A721.721.4.2";

    // -- Storage Variables --
    mapping(uint256 => mapping(address => bool)) public registeredToken;
    mapping(bytes32 => bool) public filledSwap;
    mapping(uint256 => mapping(address => address)) public swapMappingIncoming;
    mapping(uint256 => mapping(address => address)) public swapMappingOutgoing;

    /* Events from outgoing messages */
    event SwapPairRegister(
        address indexed sponsor,
        address indexed tokenAddress,
        string tokenName,
        string tokenSymbol,
        uint256 toChainId,
        uint256 feeAmount
    );

    event SwapStarted(
        address indexed tokenAddr,
        address indexed sender,
        address indexed recipient,
        uint256 dstChainId,
        uint256 tokenId,
        uint256 feeAmount
    );

    event BackwardSwapStarted(
        address indexed mirroredTokenAddr,
        address indexed sender,
        address indexed recipient,
        uint256 dstChainId,
        uint256 tokenId,
        uint256 feeAmount
    );

    /* Events from incoming messages */
    event SwapPairCreated(
        bytes32 indexed registerTxHash,
        address indexed fromTokenAddr,
        address indexed mirroredTokenAddr,
        uint256 fromChainId,
        string tokenSymbol,
        string tokenName
    );

    event SwapFilled(
        bytes32 indexed swapTxHash,
        address indexed fromTokenAddr,
        address indexed recipient,
        address mirroredTokenAddr,
        uint256 fromChainId,
        uint256 tokenId
    );

    event BackwardSwapFilled(
        bytes32 indexed swapTxHash,
        address indexed tokenAddr,
        address indexed recipient,
        uint256 fromChainId,
        uint256 tokenId
    );

    function initialize() public initializer {
        __Ownable_init();
    }

    function createSwapPair(
        bytes32 registerTxHash,
        address fromTokenAddr,
        uint256 fromChainId,
        string calldata baseURI_,
        string calldata tokenName,
        string calldata tokenSymbol
    ) public onlyOwner {
        require(
            swapMappingIncoming[fromChainId][fromTokenAddr] == address(0x0),
            _ERR721_CREATE_MIRRORED_EXISTS
        );

        MirroredERC721 mirrored = new MirroredERC721();
        mirrored.initialize(tokenName, tokenSymbol);

        if (bytes(baseURI_).length > 0) {
            mirrored.setBaseURI(baseURI_);
        }

        swapMappingIncoming[fromChainId][fromTokenAddr] = address(mirrored);
        swapMappingOutgoing[fromChainId][address(mirrored)] = fromTokenAddr;

        emit SwapPairCreated(
            registerTxHash,
            fromTokenAddr,
            address(mirrored),
            fromChainId,
            tokenSymbol,
            tokenName
        );
    }

    function registerSwapPair(address tokenAddr, uint256 chainId)
        external
        payable
    {
        require(
          tokenAddr.supportsInterface(_ERC_721_INTERFACE_ID),
          _ERR721_REGISTER_NOT_721
        );

        require(
            !registeredToken[chainId][tokenAddr],
            _ERR721_REGISTER_TOKEN_EXISTS
        );

        registeredToken[chainId][tokenAddr] = true;
        IERC721MetadataUpgradeable meta = IERC721MetadataUpgradeable(tokenAddr);
        string memory name = meta.name();
        string memory symbol = meta.symbol();

        require(bytes(name).length > 0, _ERR721_REGISTER_EMPTY_TOKEN_NAME);
        require(bytes(symbol).length > 0, _ERR721_REGISTER_EMPTY_TOKEN_SYMBOL);

        if (msg.value != 0) {
            payable(owner()).transfer(msg.value);
        }

        emit SwapPairRegister(
            msg.sender,
            tokenAddr,
            name,
            symbol,
            chainId,
            msg.value
        );
    }

    function swap(
        address tokenAddr,
        address recipient,
        uint256 tokenId,
        uint256 dstChainId
    ) external payable {
        if (msg.value != 0) {
            payable(owner()).transfer(msg.value);
        }

        // try forward swap
        if (registeredToken[dstChainId][tokenAddr]) {
            IERC721 token = IERC721(tokenAddr);
            token.safeTransferFrom(msg.sender, address(this), tokenId);

            emit SwapStarted(
                tokenAddr,
                msg.sender,
                recipient,
                dstChainId,
                tokenId,
                msg.value
            );

            return;
        }

        // try backward swap
        address dstTokenAddr = swapMappingOutgoing[dstChainId][tokenAddr];
        if (dstTokenAddr != address(0x0)) {
            IERC721Mirrored mirroredToken = IERC721Mirrored(tokenAddr);
            mirroredToken.safeTransferFrom(msg.sender, address(this), tokenId);
            mirroredToken.burn(tokenId);

            emit BackwardSwapStarted(
                tokenAddr,
                msg.sender,
                recipient,
                dstChainId,
                tokenId,
                msg.value
            );

            return;
        }

        revert(_ERR721_SWAP_NO_PAIR);
    }

    function fill(
        bytes32 swapTxHash,
        address fromTokenAddr,
        address recipient,
        uint256 fromChainId,
        uint256 tokenId,
        string calldata tokenURI
    ) public onlyOwner {
        require(!filledSwap[swapTxHash], _ERR721_FILL_ALREADY_FILLED);
        filledSwap[swapTxHash] = true;

        // fill forward swap, it means our core server will find the related mirrored token
        // and assign the value to fromTokenAddr
        address mirroredTokenAddr = swapMappingIncoming[fromChainId][
            fromTokenAddr
        ];
        if (mirroredTokenAddr != address(0x0)) {
            IERC721Mirrored mirroredToken = IERC721Mirrored(mirroredTokenAddr);
            mirroredToken.safeMint(recipient, tokenId);
            mirroredToken.setTokenURI(tokenId, tokenURI);

            emit SwapFilled(
                swapTxHash,
                fromTokenAddr,
                recipient,
                mirroredTokenAddr,
                fromChainId,
                tokenId
            );

            return;
        }

        // fill backward swap, it means that this token is the one users have been sent before
        // our server will find this token from the given mirrored token in the BackwardSwapStarted event
        // and assign the value to fromTokenAddr
        if (registeredToken[fromChainId][fromTokenAddr]) {
            IERC721 token = IERC721(fromTokenAddr);
            token.safeTransferFrom(address(this), recipient, tokenId);

            emit BackwardSwapFilled(
                swapTxHash,
                fromTokenAddr,
                recipient,
                fromChainId,
                tokenId
            );

            return;
        }

        revert(_ERR721_FILL_NO_PAIR);
    }
}
