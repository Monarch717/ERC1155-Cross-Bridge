// SPDX-License-Identifier: MIT
pragma solidity ^0.8.2;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "./MockERC1155.sol";
import "./MockERC721.sol";

contract MockFactory is Ownable {
  event ERC721Generated(
      address indexed tokenAddr
  );

  event ERC1155Generated(
      address indexed tokenAddr
  );

  address[] public erc721Tokens;
  address[] public erc1155Tokens;

  function generate721Token(
    string calldata name,
    string calldata symbol,
    string calldata baseURI,
    uint256 tokenId
  ) public onlyOwner returns (address) {
    MockERC721 token = new MockERC721(name, symbol); 
    token.setBaseURI(baseURI);
    token.safeMint(msg.sender, tokenId);
    token.setTokenURI(tokenId, string(abi.encodePacked(tokenId)));
    address tokenAddr = address(token);
    erc721Tokens.push(tokenAddr);

    emit ERC721Generated(tokenAddr);

    return tokenAddr;
  }

  function generate1155Token(
    string calldata uri,
    uint256[] memory ids,
    uint256[] memory amounts
  ) public onlyOwner returns (address) {
    MockERC1155 token = new MockERC1155(uri); 
    token.mintBatch(msg.sender, ids, amounts, "0x00");
    address tokenAddr = address(token);
    erc1155Tokens.push(tokenAddr);

    emit ERC1155Generated(tokenAddr);

    return tokenAddr;
  }
}
