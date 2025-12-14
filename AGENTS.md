Ce projet est un projet de challenge de code. L'objectif est pour l'utilisateur d'apprendre des façons de faire, d'essayer de réfléchir aux solutions éventuellement de façon astucieuse.

# ce que tu ne doit PAS faire
* résoudre les challenges en invoquant des concepts sans l'expliquer à l'utilisateur
* ne pas respecter les consignes de l'utilisateur même si elles sont fausses. Néanmoins lorsque tu as réussi à valider le challenge sur l'exemple de l'énoncé, tu peux donner des conseils à l'utilisateurs de meilleurs façons de faire les choses. Seul l'utilisateur pourra te donner son accord
* importer toi-même des librairies. Seul lo est une librairie autorisée dans ce projet

# ce que tu DOIS faire
* faire toute la glue technique très impérative peu structurante
* suivre les consignes d'architecture que l'utilisateur donnera en parallèle de l'énoncé
* accepter de suivre des consignes que tu juges mauvaises, c'est à l'utilisateur de se rendre compte qu'il te dit n'importe quoi. Pas à toi.

## workflow de travail
* l'utilisateur te copiera un énoncé en te donnant l'année et le jour, possiblement suffixé de la lettre b. Cela déterminera les 2 sous-dossiers de travail.
* le dataset étant commun entre la version A et la version b, il n'est versionné que à un seul endroit, dans la partie principale. Quand on est dans un dossier de travail principal, d'abord tu devras copier le dataset d'exemple dans le fichier txt du challenge, puis lorsque celui-ci est conforme à l'énoncé, on pourra travailler sur le dataset d'input.
* l'utilisateur se chargera lui-même de commiter le code et le dataset une fois le challenge validé. Ca te permet pour la partie B d'écraset le puzzle input avec l'exemple de l'énoncé et d'utiliser git pour revenir au puzzle input.
* pour tester lance exclusivement la commande `./run/sh année jour` , ne fait PAS de cd. 