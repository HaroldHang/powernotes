/** @type {import('tailwindcss').Config} */
const colors = require('tailwindcss/colors')
module.exports = {
    mode: 'jit',
    darkMode : 'class',
    content: [
        "./src/**/*.{js,jsx,ts,tsx}",
    ],
    theme: {
        colors: {
            transparent: 'transparent',
            current: 'currentColor',
            black: colors.black,
            white: colors.white,
            gray: colors.neutral,
            indigo: colors.indigo,
            red: colors.rose,
            green:colors.green,
            yellow: colors.amber,
            blue: colors.blue,
            purple: colors.purple,
            rose: colors.rose,
            neutral : colors.neutral,
            slate : colors.slate,
            liver: {
                DEFAULT: '#474954',
                
            },
    
        },
        fontFamily: {
            'sans': ['ui-sans-serif', 'system-ui'],
            'serif': ['ui-serif', 'Georgia'],
            'mono': ['ui-monospace', 'SFMono-Regular'],
            'display': ['Oswald'],
            'body': ['"Noto Sans","Open Sans"'],
        },
    
        extend: {
            height: {
                //sm: '8px',
                //md: '16px',
                //lg: '24px',
                //xl: '48px',
            },
            fontSize: {
                '4dxl':'2.75rem',
            },
            screens : {
                //'sd' : '468px',
                //'mlm': '500px',
                //'mld': '720px',
                //'lb' : '1024px'
            }
        },
    },
    plugins: [
        //require('flowbite/plugin')
    ],
  }