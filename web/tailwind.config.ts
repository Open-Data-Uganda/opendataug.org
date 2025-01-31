import typography from '@tailwindcss/typography';
import type { Config } from 'tailwindcss';

export default {
	content: ['./src/**/*.{html,js,svelte,ts}'],
	theme: {
		extend: {
			colors: {
				primary: '#214454',
				secondary: '#80CAEE'
			},
			screens: {
				'2xsm': '375px',
				xsm: '425px',
				'3xl': '2000px'
			},
			fontSize: {
				'title-xxl': ['44px', '55px'],
				13: '13px',
				'title-xl': ['36px', '45px'],
				'title-xl2': ['33px', '45px'],
				'title-lg': ['28px', '35px'],
				'title-md': ['24px', '30px'],
				'title-md2': ['26px', '30px'],
				'title-sm': ['20px', '26px'],
				'title-xsm': ['18px', '24px']
			},

			opacity: {
				65: '.65'
			},
			backgroundImage: {
				video: "url('../images/video/video.png')"
			},
			content: {
				'icon-copy': 'url("../images/icon/icon-copy-alt.svg")'
			},
			transitionProperty: { width: 'width', stroke: 'stroke' },
			borderWidth: {
				6: '6px'
			}
		}
	},

	plugins: [typography]
} satisfies Config;
