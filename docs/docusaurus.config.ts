import type * as Preset from "@docusaurus/preset-classic";
import type { Config } from "@docusaurus/types";
import { themes as prismThemes } from "prism-react-renderer";

const config: Config = {
  title: "Open Data Uganda | API Documentation",
  tagline: "Dinosaurs are cool",
  favicon: "img/favicon.ico",

  url: "https://docs.opendataug.org",
  baseUrl: "/",

  organizationName: "Open Data Uganda",

  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",

  trailingSlash: false,

  presets: [
    [
      "classic",
      /** @type {import('@docusaurus/preset-classic').Options} */
      {
        docs: {
          sidebarPath: "./sidebars.ts",
          routeBasePath: "/",
        },
        pages: false,
        blog: false,
        theme: {
          customCss: "./src/css/custom.css",
        },
      } satisfies Preset.Options,
    ],
  ],

  themeConfig: {
    image: "img/docusaurus-social-card.jpg",
    colorMode: {
      defaultMode: "light",
      disableSwitch: true,
      respectPrefersColorScheme: false,
    },
    navbar: {
      logo: {
        alt: "Open Data Uganda Logo",
        src: "img/logo.png",
      },
      items: [
        {
          type: "html",
          value: "Documentation",
          position: "left",
        },
        {
          type: "search",
          position: "left",
        },
      ],
    },

    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.duotoneLight,
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
