import { createI18n } from 'vue-i18n';

export const i18n = createI18n({
    locale: 'fr',
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
                btn_label: "Surprise me!"
            },
            config: {
                language: "Language",
                nbrPkmn: "Number of Pokémons to show",
                randomizerLanguage: "Starter name language displayed",
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
                btn_label: "Surprends moi !"
            },
            config: {
                language: "Langue",
                nbrPkmn: "Nombre de Pokémons à afficher",
                randomizerLanguage: "Langue utilisée pour les starters",
            }
        }
    }
})