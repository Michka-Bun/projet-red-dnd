<div align="center">

<h1>⚔️ R-E-D&amp;D</h1>

<a href="https://go.dev/">
  <img alt="Go" src="https://img.shields.io/badge/Go-latest-00ADD8?logo=go&logoColor=white" />
</a>

<p><em>Un RPG minimaliste en ligne de commande, écrit en Go, pour apprendre et pratiquer la programmation.</em></p>

<p>
  <kbd>CLI</kbd>
  <kbd>Turn-Based</kbd>
  <kbd>Single-Player</kbd>
  <kbd>No Third-Party Modules</kbd>
</p>

</div>

---

<details>
  <summary><strong>📑 Table of Contents (cliquer pour dérouler)</strong></summary>

- [Description](#-description)
  - [Gameplay](#-gameplay-vue-densemble)
- [Fonctionnalités de gameplay](#-fonctionnalités-de-gameplay)
- [Stack](#-stack)
- [Installation / Lancer le jeu](#-installation--lancer-le-jeu)
- [Structure du projet](#-structure-du-projet)
- [Contributeurs](#-contributeurs)

</details>

---

## 📝 Description
> **R-E-D&amp;D** est un **RPG minimaliste** en **ligne de commande**.  
> Il sert de terrain d’entraînement pour consolider ses bases de programmation tout en proposant une boucle de jeu claire :  
> **création de personnage**, **exploration par menus**, **inventaire/équipement**, et **combats au tour par tour** avec **effets d’état**.

---

### 🎮 Gameplay (vue d’ensemble)
- **Création de personnage** : nom + **classe** parmi *Warrior*, *Mage*, *Viking* et *Archer*.  
- **Menu principal** : fiche perso, **Inventory**, **Shop**, **Blacksmith**, **Fight**.  
- **Combat au tour par tour** : ennemis avec **faiblesses/résistances**, gestion **HP/Mana**, **compétences** et **objets** (potions), **Poison** sur plusieurs tours, etc.  
- **Progression** : **LevelUp**, statistiques, **équipement** (bonus via `EquipmentStats`), optimisation via forgeron et achats.

> Aperçu (extrait CLI)
> ```text
> --------------------
> Menu:
> 1. Character info
> 2. Inventory
> 3. Shop
> 4. Blacksmith
> 5. Fight
> 6. Exit
> Choose an option:
> ```

---

## ✨ Fonctionnalités de gameplay
<div>

### Classes & faiblesses
<table>
  <thead>
    <tr>
      <th align="left">Classe</th>
      <th align="left">Style</th>
      <th align="left">Faiblesse thématique</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>🛡️ Warrior</td>
      <td>Frontline, tanking et coups directs</td>
      <td><kbd>Weak vs Piaf</kbd></td>
    </tr>
    <tr>
      <td>🧙 Mage</td>
      <td>Dégâts magiques, gestion de Mana</td>
      <td><kbd>Weak vs Devourer</kbd></td>
    </tr>
    <tr>
      <td>🪓 Viking</td>
      <td>Brutal, HP élevés, impact soutenu</td>
      <td><kbd>Weak vs Piaf</kbd></td>
    </tr>
    <tr>
      <td>🏹 Archer</td>
      <td>DPS à distance, gestion des ressources</td>
      <td><kbd>Weak vs Devourer</kbd></td>
    </tr>
  </tbody>
</table>

### Système de combat
- **Tour par tour** : choisir entre **attaques**, **compétences**, **objets**, **équipement**.
- **Effets d’état** : ex. **Poison** (dégâts récurrents sur N tours).  
- **Barres** : `HPBar` / `ManaBar` (affichage ANSI coloré).  
- **Checks de mort** : `IsDead` et messages d’alerte clairs.

### Inventaire & objets
- **Potions** : Vie (**Health**), **Mana**, **Poison**.  
- **Gestion** : `AddItem`, `RemoveItem`, affichage soigné (`FormatInventory`).
- **Utilisation contextuelle** : effets appliqués au joueur et/ou au monstre selon l’objet.

### Équipement & progression
- **Équiper/Déséquiper** : `EquipItem`, `UnequipItem`.  
- **Bonus** : via `EquipmentStats` (ex. HPmax/Manamax, dégâts).  
- **Économie** : **Shop** (achat) et **Blacksmith** (améliorations).

</div>

---

## 🛠️ Stack
- **Langage** : Go (**dernière version**)
- **Standard library** : `fmt`, `os`, `strings`, `math`, `sort`  

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

<div align="center">
👥 Contributeurs

<strong>Thibaud SELLIER</strong> • <strong>Michel LEVINE</strong>

</div>
