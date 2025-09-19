# R-E-D&D

[![Go](https://img.shields.io/badge/Go-latest-00ADD8?logo=go&logoColor=white)](https://go.dev/)

## Table of Contents
- [Description](#-description)
  - [Gameplay](#-gameplay-vue-densemble)
- [Fonctionnalités de gameplay](#-fonctionnalités-de-gameplay)
- [Stack](#-stack)
- [Installation / Lancer le jeu](#-installation--lancer-le-jeu)
- [Structure du projet](#-structure-du-projet)
- [Contributeurs](#-contributeurs)

---

## 📝 Description
**R-E-D&D** est un **jeu RPG minimaliste en ligne de commande** écrit en Go. Il sert de terrain d’entraînement pour consolider ses bases de programmation tout en proposant une boucle de jeu claire : création de personnage, exploration menu-driven, gestion d’inventaire/équipement, et combats au tour par tour avec effets d’état.

---

### 🎮 Gameplay (vue d’ensemble)
- **Création de personnage** : choix du nom et de la **classe** parmi *Warrior*, *Mage*, *Viking* et *Archer*.  
- **Menu principal** : accéder à la fiche perso, gérer l’inventaire, visiter le **Shop** (achat) et le **Blacksmith** (améliorations/équipement), ou lancer des **combats**.  
- **Combats au tour par tour** : affrontements textuels avec **faiblesses**/résistances selon les types d’ennemis et la classe du joueur. Gestion des **HP**, **Mana**, **compétences**, **objets** (potions), et **effets d’état** (ex. **Poison** qui inflige des dégâts sur plusieurs tours).  
- **Progression** : montée de niveau (*LevelUp*), amélioration des statistiques, gestion d’**équipement** (bonus via `EquipmentStats`), et optimisation du build via le forgeron et les achats.

---

## ✨ Fonctionnalités de gameplay
- **Classes jouables** : *Warrior*, *Mage*, *Viking*, *Archer* — chacune interagit différemment avec certains ennemis (faiblesses/avantages thématiques).
- **Bestiaire & faiblesses** : plusieurs **monstres** avec comportements/valeurs distincts et **faiblesses de type** (pénalités/bonus en fonction de la classe ou de la nature de l’ennemi).
- **Système de combat** :
  - **Tour par tour** avec choix entre **attaques**, **compétences** (consomment du **Mana**), **objets** et **équipement**.
  - **Effets d’état** : **Poison** (dégâts récurrents sur N tours), autres variations selon compétences/ennemis.
  - **Mort & checks** : gestion de l’état de mort via `IsDead`, affichages d’alertes, barres **HP/Mana** avec mise en forme ANSI.
- **Inventaire & objets** :
  - **Potions** : **Health potion** (régénère des HP), **Mana potion** (régénère du Mana), **Poison potion** (appliquée à l’ennemi ou auto-effet si mal utilisée).
  - Ajout/consommation (`AddItem` / `RemoveItem`) et affichage propre (`FormatInventory`).
- **Équipement & stats** :
  - **Équiper/Déséquiper** avec `EquipItem`/`UnequipItem` et application des **bonus** via `EquipmentStats`.
  - Progression des **HPmax/Manamax**, **dégâts**, etc., via **niveaux** et **forge**.
- **Économie & services** :
  - **Shop** : achat de potions/équipement.
  - **Blacksmith** : amélioration/gestion de l’équipement.
- **Interface CLI** :
  - **Menu clair** (fiche perso, inventaire, shop, blacksmith, combat).
  - **Barres de vie/mana** (`HPBar`, `ManaBar`) et messages colorés pour la lisibilité.

---

## 🛠️ Stack
- **Langage** : Go (dernière version)
- **Bibliothèques standard** : `fmt`, `os`, `strings`, `math`, `sort`

---

## 🚀 Installation / Lancer le jeu
```bash
git clone https://github.com/Michka-Bun/projet-red-dnd.git
cd projet-red-dnd/src
go run main/main.go
```

---

## 📂 Structure du projet
```bash
src/
├─ main/main.go          # Point d’entrée (boucle de jeu, navigation menus)
├─ character.go          # Définition Player, stats, équipement, inventaire, level-up
├─ fight.go              # Logique de combat au tour par tour, actions/skills, checks de mort
├─ menu.go               # Menus: fiche perso, inventaire, shop, blacksmith, combat
├─ monsters.go           # Définition des monstres + faiblesses/paramètres
├─ skills.go             # Compétences (coûts mana, effets, calculs de dégâts)
└─ useitems.go           # Objets (potions HP/Mana/Poison), effets et messages utilisateur
```

---

## 👥 Contributeurs
- Thibaud SELLIER
- Michel LEVINE