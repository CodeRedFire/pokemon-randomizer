import { createI18n } from 'vue-i18n';

export const i18n = createI18n({
    locale: 'fr',
    fallbackLocale: 'en',
    messages: {
        en: {
            languages: {
                jaHrkt: "Japanese - Hiragana/Katakana",
                roomaji: "Japanese - Roomaji",
                ko: "Korean",
                zhHant: "Traditional Chinese",
                fr: "French",
                en: "English",
                de: "Deutsch",
                es: "Spanish",
                it: "Italian",
                ja: "Japanese",
                zhHans: "Simplified Chinese"
            }
        },
        fr: {
            languages: {
                jaHrkt: "Japonais - Hiragana/Katakana",
                roomaji: "Japonais - Roomaji",
                ko: "Coréen",
                zhHant: "Chinois traditionnel",
                fr: "Français",
                en: "Anglais",
                de: "Allemand",
                es: "Espagnol",
                it: "Italien",
                ja: "Japonais",
                zhHans: "Chinois simplifié"
            }
        }
    }
})