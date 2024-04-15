# Projet Hiver 2024

Ce projet est une exploration de WASM et Kubernetes dans le cadre d'un jeu multijoueur en ligne.
Le jeu est un platformer multijoueur en temps réel facilement déployable localement à l'aide de `Docker` ou `Kubernetes`.

Certaines limitations sont inévitables dû à la nature de WASM qui est lui-même un projet en phase expérimentale.

### Pour un bref moment, le jeu sera disponible en ligne à l'aide d'un déploiement avec Kubernetes

Vous pouvez accéder au jeu en suivant le lien suivant: [http://online-client.cleroux.dev](http://online-client.cleroux.dev)

## Requirements

Pour rouler le projet plus facilement, les applications sont offertes sous forme de docker container
disponible dans le GHCR. Pour déployer le projet localement, il faudra alors au minimum une installation
fonctionnelle de docker.

Il est également possible de build le projet localement. Pour faciliter l'installation des dépendances,
un fichier `flake.nix` est fournit dans le projet. Pour l'utiliser, il faudra installer `nix` avec
l'option `flake enabled`. Je recommande fortement [nix-installer](https://github.com/DeterminateSystems/nix-installer)
puisqu'il installe `nix` avec toutes les options nécessaires "out-of-the-box" sur toutes les grandes plateformes.

_TLDR: Plus de détails sont disponibles dans le répertoires de [nix-installer](https://github.com/DeterminateSystems/nix-installer),
mais cette commande devrait installer `nix` et voud serez en mesure de rouler le projet localement:_

```bash
curl --proto '=https' --tlsv1.2 -sSf -L https://install.determinate.systems/nix | sh -s -- install
```

## Build

Si vous déployez avec `docker compose`, les dépendances du projet sont inclues dans l'image de build. Vous n'aurez donc pas besoin d'installer
nix. Toutefois, si vous voulez tester un cluster Kubernetes localement, vous aurez besoin d'activer l'environnement de développement pour avoir
accès à [kind](https://kind.sigs.k8s.io/) ou vous pouvez l'installer vous même.

De plus, pour rouler les tests avec la commande `go test`, vous aurez également besoin de l'environnement de développement nix ou d'installer toutes
les dépendances manuellement.

### Activer l'environnement de développement

Pour activer l'environnement de développement, vous n'avez qu'à faire la commande suivante dans le `PROJECT_ROOT`:

```bash
nix develop
```

Cette commande va initialiser un environnement temporaire et installer toutes les dépendances du projet dans cet environnement.
Cela ne fera aucune modification au système global.

Il n'est pas nécessaire de build le projet localement pour l'utiliser, toutefois si vous voulez changer les paramètres de l'application,
vous devrez absolument build localement. Comme le projet est un module WASM, il est malheureusement impossible de passer des `flags` ou des `ENV VARS` à
l'application. Effectivement, Go compile des modules WASM comme un runtime isolé qui ne communique pas avec le host.

Si vous désirez modifier les paramètres du jeu, vous pouvez changer le fichier de configuration `internal/game/config.yaml` puis recompiler.

### Docker Compose

Si vous prévoyez déployer localement avec `docker compose`, utilisez cette commande:

```bash
docker compose build client server
```

### Kubernetes

Si vous prévoyez déployer localement avec `Kubernetes/Kind`, utilisez cette commande:

```bash
docker compose build kind-client kind-server
```

## Usage

Si vous n'avez pas l'intention d'installer `nix`, vous pourrez déployer l'application avec les instructions pour le déploiement `docker compose`.

### Docker Compose

Pour déployer le serveur de jeu localement, utiliser cette commande:

```bash
docker compose up client server
```

Le client sera accessible à l'adresse `http://localhost:8080`.

Le serveur sera accessible à l'adresse `http://localhost:8888`.

### Kubernetes

Comme déployer une application sur `Kubernetes` est naturellement complexe, un script automatisé est fournit dans le fichier `./scripts/startup.sh`.
Veuillez notez qu'il est nécessaire d'avoir `nix` d'installer sur votre machine pour utilser ce script. Pour déployer le cluster Kubernetes, utilisez cette commande:

```bash
./scripts/startup.sh
```

Le client sera accessible à l'adresse `http://localhost/client/`.

Le serveur sera accessible à l'adresse `http://localhost/server/`.

_**Note: N'oubliez pas le '/' à la fin du URL. Vous devez absoluement l'inclure pour le bon fonctionnement du reverse-proxy dans le cluster.**_
