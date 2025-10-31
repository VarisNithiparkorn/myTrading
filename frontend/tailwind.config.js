
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    // ... paths to your components
  ],
  theme: {
    extend: {
      fontFamily: {
        'custom-font': ['var(--font-mycustom)', 'sans-serif'],
      },
    },
  },
  plugins: [],
};