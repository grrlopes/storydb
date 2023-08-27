# Changelog

Important changes will be written here.

## 0.5.0 - 27-08-2023

### Fixed
- unselected row filled in terminal input
- sync screen responsiveness

### Added
- animation spin in sync screen

### Changed
- string key to mapping key in sync screen
- sync screen message status

### Removed
- progress bar
- usecase file history length

## 0.4.0-beta-0 - 19-08-2023

### Fixed
- Selection hover on list view
- Paginator on screen finder is being update 
- Home screen get update to fisrt record after
open finder textinput

### Added
- Fulltext search for command field db
- Helper keymap implemented in every screen

### Changed
- All string keymap to constant component

### Removed
- External lib that fill commandline 

## 0.3.0-alpha-0 - 03-08-2023

### Fixed
- Responsiveness is properly working.
- Finder search is being hidden.

### Added
- Sql "like" into search finder repository.
- Search finder count repository

## 0.2.0-alpha-1 - 14-07-2023

### Fixed
- Pagination is working properly based on history line
- Selection btn at syncScreen dialog

### Added
- Dialog style btn
- Auto create storydb folder
- Path for store database at $HOME/.local/share/storydb

### Changed
- How the app load the .bash_history

### Removed
- SqliteErr constant from repository ``#0c40006``

## 0.2.0-alpha - 06-07-2023

### Added
- Usecase Load file bash history file
- Custom pagination
- Sql insert repository to sink history
- Sql cound and pagination repository
- Moving hotkey ctrl+u and g

### Changed
- Entity listpanel model was replaced by custom struct  

### Removed
- Listpanel component

## 0.1.0-alpha - 21-05-2023

### Added

- Initial project scaffolder
