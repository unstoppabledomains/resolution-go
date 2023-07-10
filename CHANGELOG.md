## v3.1.0 - 2023-07-10

Add GetAddr, ReverseOf API

## v3.0.0 - 2023-04-05

Remove default RPC urls, Allow initialized with UD partner key

## 2.3.2

- Update supported keys in `keys.go`

## 2.3.1

- Update config to use Infura

## 2.3.0

- Deprecate `AllRecords` method

## 2.2.0

- Use default networks for UnsBuilder

## 2.1.0

- Update config to use Alchemy
- Update tests from Rinkeby to Goerli
- Update resolver-keys

## 2.0.0

- L2 support
- Add Locations method

## 1.1.2

- Add Namehash method

## 1.1.1

- Add UNS support
- Replaces resolution.NewCnsBuilder() with resolution.NewUnsBuilder()
- Add UNS testnet support with UnsBuilder.SetEthereumNetwork() method
- Add Unhash method
- Add TokenURIMetadata getter
- isSupportedDomain() now makes network call

## 1.0.0

- Initial release
