// SPDX-License-Identifier: MIT
pragma solidity ^0.8.2;

import "@openzeppelin/contracts-upgradeable/token/ERC1155/extensions/IERC1155MetadataURIUpgradeable.sol";

interface IERC1155Mirrored is IERC1155MetadataURIUpgradeable {
  function mint(address account, uint256 id, uint256 amount, bytes memory data) external;
  function mintBatch(address to, uint256[] memory ids, uint256[] memory amounts, bytes memory data) external;
  function burn(address account, uint256 id, uint256 value) external;
  function burnBatch(address account, uint256[] memory ids, uint256[] memory values) external;
}
