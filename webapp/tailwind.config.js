module.exports = {
  purge: ['./src/**/*.{js,jsx,ts,tsx}', './public/index.html'],
  theme: {
    extend: {
      backgroundColor: {
        'futuristic-black': '#0F0F15',
        'futuristic-blue': '#0F6FFF',
      },
      textColor: {
        'futuristic-blue': '#0F6FFF',
      },
    },
  },
  variants: {
    extend: {},
  },
  plugins: [require('daisyui')],
};
