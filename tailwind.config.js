/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./web/**/*.{html,gohtml}'],
  darkMode: 'class',
  theme: {
    screens: {
      sm: '320px',
      md: '768px',
      lg: '1024px',
      xl: '1280px',
    },
    extend: {},
  },
  plugins: [],
}

