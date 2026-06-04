# Changelog

## [1.3.1](https://github.com/nednella/bootstrap.sh/compare/v1.3.0...v1.3.1) (2026-06-04)


### Bug Fixes

* restrict the banner to job commands ([ef35f28](https://github.com/nednella/bootstrap.sh/commit/ef35f28a65471f8b475f6586531a29e632da2bdb))

## [1.3.0](https://github.com/nednella/bootstrap.sh/compare/v1.2.0...v1.3.0) (2026-06-04)


### Features

* split local repository updates into dedicated sync command ([40aa46e](https://github.com/nednella/bootstrap.sh/commit/40aa46e14dbbffb674467d15d51f2bcabbffbc4e))


### Bug Fixes

* add context to malformed settings parse error ([c710b0a](https://github.com/nednella/bootstrap.sh/commit/c710b0a9fe81b8b8dde119a51866b64db2146ea8))
* add context to repository clone failure ([2ae3c18](https://github.com/nednella/bootstrap.sh/commit/2ae3c18abf29720ab12c0f24cb88aab231a58bf1))
* clean up the staged binary on a failed update ([6227802](https://github.com/nednella/bootstrap.sh/commit/62278020b04b5b2ee0ef0c45ead41661650ad8c5))
* guard against an invalid release tag ([41d10b8](https://github.com/nednella/bootstrap.sh/commit/41d10b8756009b63b6a90c613e9c119c681f79b7))
* preserve path structure in dotfiles backups ([768cde9](https://github.com/nednella/bootstrap.sh/commit/768cde96bfbb54b5f284498a11daebb7a4798678))
* rebase with autostash during sync ([9a0a99e](https://github.com/nednella/bootstrap.sh/commit/9a0a99e269d33aa23e6fd165ecfa508e241b0d4e))
* reject incomplete macOS settings ([8e26399](https://github.com/nednella/bootstrap.sh/commit/8e26399fc4ba001b1d4ac89832b345cd076c6f40))
* skip self-update on development builds ([d9b0fc9](https://github.com/nednella/bootstrap.sh/commit/d9b0fc97e58efa86a33cb3b33b018986e62624d1))
* strip quarantine xattr on self-update ([3a9234f](https://github.com/nednella/bootstrap.sh/commit/3a9234f0c0e09fe77a4b349a016d23ef43579a47))

## [1.2.0](https://github.com/nednella/bootstrap.sh/compare/v1.1.1...v1.2.0) (2026-06-04)


### Features

* externalise macOS defaults to runtime settings.yaml ([d69fe84](https://github.com/nednella/bootstrap.sh/commit/d69fe84a78f4bc42e2037a3945e2b9964eb54e60))

## [1.1.1](https://github.com/nednella/bootstrap.sh/compare/v1.1.0...v1.1.1) (2026-06-03)


### Bug Fixes

* banner spacing ([a9d9669](https://github.com/nednella/bootstrap.sh/commit/a9d96698e468fd63059b6ff50bd72cfa601ae126))

## [1.1.0](https://github.com/nednella/bootstrap.sh/compare/v1.0.0...v1.1.0) (2026-06-03)


### Features

* add update job ([33c84f8](https://github.com/nednella/bootstrap.sh/commit/33c84f89cb3d8ae5bc8517eb55c8f7fbad496ae7))
* track the binary's version ([93511a0](https://github.com/nednella/bootstrap.sh/commit/93511a02f19940e101ad9b8969e72e22c60128dd))


### Bug Fixes

* preflight dry-run prints the real command ([1e8209f](https://github.com/nednella/bootstrap.sh/commit/1e8209f86cb5f304a0b4f8765e46bf4206f85283))

## 1.0.0 (2026-06-03)


### Features

* add --dry-run flag ([c5fe1cf](https://github.com/nednella/bootstrap.sh/commit/c5fe1cf77c47d6b85044c427fdd0a89a349ce965))
* add cobra root command ([a60991d](https://github.com/nednella/bootstrap.sh/commit/a60991d083768193e1de68d4a622aefa23368058))
* add config loader ([8de044f](https://github.com/nednella/bootstrap.sh/commit/8de044f796f24de17966c9cd1076b618a8361ac8))
* add dotfiles command skeleton ([5762a0d](https://github.com/nednella/bootstrap.sh/commit/5762a0da8c4d16924fe777da8af89bd825db2a25))
* add dotfiles job ([45606ed](https://github.com/nednella/bootstrap.sh/commit/45606ed8cf27a7b0e8490d239d7ba655d5c35433))
* add install command skeleton ([ef0370b](https://github.com/nednella/bootstrap.sh/commit/ef0370b48dc3a5e01d82f724a687f7db785627d8))
* add install job ([34a7e54](https://github.com/nednella/bootstrap.sh/commit/34a7e54e649995f767e3a1d0cff27ae1b35f18e0))
* add install.sh ([f624d97](https://github.com/nednella/bootstrap.sh/commit/f624d97341d3d392fcdab3d83a9ac1bc4aeaf71f))
* add macos command skeleton ([0379f0c](https://github.com/nednella/bootstrap.sh/commit/0379f0c33bb90aebf745e51c227f140239da6ee5))
* add macos job ([f490259](https://github.com/nednella/bootstrap.sh/commit/f49025983bdede55162ac5f854b9c88cdc92a3a3))
* add shell-out runner ([4cdc4c4](https://github.com/nednella/bootstrap.sh/commit/4cdc4c463f8c8f5f6d4aa2af98ddbba598c70144))
* add ui banner ([65f138a](https://github.com/nednella/bootstrap.sh/commit/65f138a0b0fab63bea9a92c1c40c3709dc428d0d))
* add ui logging ([72763d6](https://github.com/nednella/bootstrap.sh/commit/72763d68e052cc2b71d95e01efe7b263bcfc5f9d))
* add update command skeleton ([e3fe924](https://github.com/nednella/bootstrap.sh/commit/e3fe92461ed1cd20ab67bd6ae282b9e1d75f2531))
* group jobs in help output ([e9682d7](https://github.com/nednella/bootstrap.sh/commit/e9682d709d05f316dc2a37ba4950e0f1e56e63bc))
* init CLI tool project ([0b64528](https://github.com/nednella/bootstrap.sh/commit/0b6452834bbe4100347a84cf4229945acfe6b29e))
* integrate ui ([96f695c](https://github.com/nednella/bootstrap.sh/commit/96f695c9f082e61a1af33f5e87471bc3c141e623))
* run preflight before job commands ([6268ae3](https://github.com/nednella/bootstrap.sh/commit/6268ae3da5f9d64e0992c7751d30a99df704afe0))
