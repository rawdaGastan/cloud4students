/**
 * plugins/vuetify.js
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

import { createVuetify } from 'vuetify'

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
    theme: {
        themes: {
            light: {
                colors: {
                    primary: '#1867C0',
                    secondary: '#5CBBF6',
                    background: '#D8F2FA',
                    accent: '#FFFFFF',
                },
            },
            dark: {
                colors: {
                    primary: '#1982b1',
                    secondary: '#5CBBF6',
                    background: '#333'
                },
            },
        },
    },
})