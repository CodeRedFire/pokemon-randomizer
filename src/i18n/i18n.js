import { createI18n } from 'vue-i18n';

export const i18n = createI18n({
    locale: 'en',
    fallbackLocale: 'en',
    messages: {
        en: {
            languages: {
                jaHrkt: "Japanese - Hiragana/Katakana",
                roomaji: "Japanese - Roomaji",
                koLatin: "Korean (Latin alphabet)",
                koHangul: "Korean (Hangul alphabet)",
                zhHant: "Traditional Chinese",
                fr: "French",
                en: "English",
                de: "Deutsch",
                es: "Spanish",
                it: "Italian",
                ja: "Japanese",
                zhHans: "Simplified Chinese",
                zhLatin: "Chinese (Latin alphabet)"
            },
            randomizer: {
                btn_label: "Surprise me!",
                selection_tooltip: "Select the pokémon",
                team: {
                    reset_tooltip: "Reset the selected team"
                }
            },
            config: {
                general: {
                    title: "General",
                    language: "Language",
                },
                battleship: {
                    title: "Battleship",
                    column: "Number of columns",
                    row: "Number of rows"
                },
                randomizer: {
                    title: "Randomizer",
                    nbrPkmn: "Number of Pokémons to show",
                    randomizerLanguage: "Starter name language displayed",
                    displayedInformations: "Displayed informations",
                    all: "All",
                    nameOnly: "Name only",
                    bestStat: "Best stat",
                    worstStat: "Worst stat",
                    type: "Type",
                    pokedexNumber: "Pokedex number",
                }
            },
            menu: {
                randomizer: "Pokémon Randomizer",
                battleship: "Battleship",
                title:"Choose your game mode"
            },
            battleship: {
                regenerate: "Regenerate grid",
            }
        },
        fr: {
            languages: {
                jaHrkt: "Japonais - Hiragana/Katakana",
                roomaji: "Japonais - Roomaji",
                koLatin: "Coréen (alphabet latin)",
                koHangul: "Coréen (alphabet hangul)",
                zhHant: "Chinois traditionnel",
                fr: "Français",
                en: "Anglais",
                de: "Allemand",
                es: "Espagnol",
                it: "Italien",
                ja: "Japonais",
                zhHans: "Chinois simplifié",
                zhLatin: "Chinois (alphabet latin)"
            },
            randomizer: {
                btn_label: "Surprends moi !",
                selection_tooltip: "Sélectionne le pokémon",
                team: {
                    reset_tooltip: "Réinitialise l'équipe sélectionnée"
                }
            },
            config: {
                general: {
                    title: "Général",
                    language: "Langue",
                },
                battleship: {
                  title: "Bataille Navale",
                  column: "Nombre de colonnes",
                  row: "Nombre de lignes"
                },
                randomizer: {
                    title: "Randomizer",
                    nbrPkmn: "Nombre de Pokémons à afficher",
                    randomizerLanguage: "Langue utilisée pour les starters",
                    displayedInformations: "Informations affichées",
                    all: "Tout",
                    nameOnly: "Nom uniquement",
                    bestStat: "Meilleure statistique",
                    worstStat: "Pire statistique",
                    type: "Type",
                    pokedexNumber: "Numéro de Pokédex",
                }
            },
            menu: {
                randomizer: "Randomizer Pokémon",
                battleship: "Bataille Navale",
                title:"Choisissez votre mode de jeu"
            },
            battleship: {
                regenerate: "Regénérer la grille",
            }
        }
    }
})