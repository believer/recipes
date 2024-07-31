/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		"./views/**/*.html",
		"./views/**/*.templ",
		"./model/*.go",
		"./styles/styles.css",
	],
	theme: {
		extend: {},
	},
	plugins: [],
};
