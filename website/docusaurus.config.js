module.exports = {
  title: 'Opctl',
  tagline: 'Automate operating your project using containers as building blocks.',
  url: 'https://opctl.io',
  baseUrl: '/opctl/',
  favicon: 'img/favicon.ico',
  organizationName: 'apexskier',
  projectName: 'opctl',
  trailingSlash: false,
  themeConfig: {
    algolia: {
      appId: 'E19H3NL09D',
      apiKey: '247689f43002ed0c3902fbd034630a7a',
      indexName: 'opctl'
    },
    navbar: {
      title: 'Opctl',
      logo: {
        alt: 'opctl Logo',
        src: 'img/logo.svg',
      },
      items: [
        { to: 'docs/introduction', label: 'Docs', position: 'left' },
        {
          href: 'https://github.com/orgs/opspec-pkgs/repositories',
          label: 'Packages',
          position: 'left',
        },
        {
          href: 'https://join.slack.com/t/opctl/shared_invite/zt-51zodvjn-Ul_UXfkhqYLWZPQTvNPp5w',
          label: 'Slack',
          position: 'right',
        },
        {
          href: 'https://github.com/opctl/opctl',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Docs',
              to: 'docs/introduction',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Slack',
              href: 'https://join.slack.com/t/opctl/shared_invite/zt-51zodvjn-Ul_UXfkhqYLWZPQTvNPp5w',
            },
          ],
        },
        {
          title: 'Social',
          items: [
            {
              label: 'Github',
              href: 'https://github.com/opctl/opctl'
            }
          ]
        }
      ],
      copyright: `Copyright © ${new Date().getFullYear()} opctl.io`,
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          editUrl: "https://github.com/opctl/opctl/edit/main/website/",
          sidebarPath: require.resolve('./sidebars.js'),
          // Equivalent to `enableUpdateBy`.
          showLastUpdateAuthor: true,
          // Equivalent to `enableUpdateTime`.
          showLastUpdateTime: true,
        },
        googleAnalytics: {
          trackingID: 'UA-94109316-1',
        }
      },
    ],
  ],
};
