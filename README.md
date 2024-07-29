
<div align="center">
<img src="https://github.com/Mugen-Builders/.github/assets/153661799/7ed08d4c-89f4-4bde-a635-0b332affbd5d" width="150" height="150">
</div>
<br>
<div align="center">
<i>DeVolt is a descentralized solution focused on providing the eletricity needed for eletric cars.</i>
</div>
<div align="center">
<b>With monetary incentives, logistical facilitations, more accessible stations, and an open, fully transparent market.<b>
</div>
<br>
<div align="center">
</a> <a href="https://mugen-builders.github.io/devolt/">
<img src="https://img.shields.io/badge/devolt-Website-85EA51"/> </a>
<a href="https://x.com/ddevolt/">
<img src="https://img.shields.io/twitter/follow/ddevolt?style=social"/>
</a> <a href="https://mugen-builders.github.io/devolt/">
<img src="https://img.shields.io/badge/docs-Website-yellow"/> </a>

<a href="https://docs.cartesi.io/cartesi-rollups/">![Static Badge](https://img.shields.io/badge/cartesi-1.3.0-5bd1d7)</a>
<a href="https://docs.cartesi.io/cartesi-rollups/1.3/quickstart/">![Static Badge](https://img.shields.io/badge/cartesi--cli-0.15.0-5bd1d7)</a>
<a href="https://pkg.go.dev/github.com/calindra/nonodo">![Static Badge](https://img.shields.io/badge/nonodo-1.1.1-blue)</a>
<a href="https://pkg.go.dev/github.com/gligneul/rollmelette">![Static Badge](https://img.shields.io/badge/rollmelette-0.1.1-yellow)</a>
<a href="https://book.getfoundry.sh/getting-started/installation">![Static Badge](https://img.shields.io/badge/foundry-0.2.0-red)</a>
<a href="https://pkg.go.dev/gorm.io/driver/sqlite@v1.5.6">![Static Badge](https://img.shields.io/badge/sqlite-1.5.6-blue)</a>
<a href="https://pkg.go.dev/gorm.io/gorm@v1.25.10">![Static Badge](https://img.shields.io/badge/gorm-1.25.10-blue)</a>
<a href="https://pkg.go.dev/github.com/google/wire@v0.6.0">![Static Badge](https://img.shields.io/badge/wire-0.6.0-blue)</a>
</div>

## 📚 Technical Vision:
This project was built using Golang as the main language and [SQLite](https://www.sqlite.org/) to store the application state, along with the ORM [Gorm](https://gorm.io/). Additionally, this project was built following the [golang-standards](https://github.com/golang-standards/project-layout) [^1], and from an architectural perspective, principles of [hexagonal architecture](https://netflixtechblog.com/ready-for-changes-with-hexagonal-architecture-b315ec967749) [^2] were implemented, such as dependency injection, using the [Wire](https://github.com/google/wire) package for automatic initialization. From a technical standpoint, this choice of architecture and technologies was made possible because we are building this application using the Cartesi infrastructure.

## 🏎️ Running:
### Local node:
- Build the application:

```bash
$ make build
```

- Run locally:

```bash
$ cartesi run
```

### A validator node on Fly.io:
- Build the validator node image
```bash
$ make build
```
- After that, you can follow the [tutorial](https://docs.cartesi.io/cartesi-rollups/1.3/deployment/self-hosted/#hosting-on-flyio) and after creating the necessary infrastructure to host your node, you can use the Docker image generated in the previous step called `validator:latest`

### Application Test:
- To run the complete test suite, run the command below:

```bash
$ make test
```

- To see the test coverage in the application, run the command below:
```bash
$ make coverage
```

## 🌐 Deployed Application:
- Node Public URL: https://devolt.fly.dev/
- Application contract address (Arbitrum Sepolia Network): [0xdDa19ea9b093Ad3a4A4DBA861EDFc20c1e1aC601](https://sepolia.arbiscan.io/address/0xdda19ea9b093ad3a4a4dba861edfc20c1e1ac601)

[^1]: The folder structure chosen for this project is in line with the conventions and standards used by the Golang developer community.

[^2]: The entities, repositories, and use cases are in accordance with the standards provided for hexagonal architecture.