// @ts-check
import { defineConfig } from 'astro/config';
import starlight from '@astrojs/starlight';

// https://astro.build/config
export default defineConfig({
	site: 'https://junara.github.io',
	base: '/jmactts',
	integrations: [
		starlight({
			title: 'jmactts',
			social: [{ icon: 'github', label: 'GitHub', href: 'https://github.com/junara/jmactts' }],
			defaultLocale: 'ja',
			locales: {
				ja: { label: '日本語', lang: 'ja' },
				en: { label: 'English', lang: 'en' },
			},
			sidebar: [
				{
					label: 'はじめに',
					translations: { en: 'Getting Started' },
					items: [
						{ label: 'インストール', slug: 'getting-started/installation', translations: { en: 'Installation' } },
						{ label: 'クイックスタート', slug: 'getting-started/quickstart', translations: { en: 'Quick Start' } },
					],
				},
				{
					label: '使い方',
					translations: { en: 'Usage' },
					items: [
						{ label: '入力ソース', slug: 'usage/input', translations: { en: 'Input Sources' } },
						{ label: 'ボイス選択', slug: 'usage/voice', translations: { en: 'Voice Selection' } },
						{ label: '出力フォーマット', slug: 'usage/output', translations: { en: 'Output Formats' } },
						{ label: '長文と中断', slug: 'usage/long-text', translations: { en: 'Long Text & Interrupt' } },
					],
				},
				{
					label: 'リファレンス',
					translations: { en: 'Reference' },
					items: [
						{ label: 'フラグ一覧', slug: 'reference/flags', translations: { en: 'Flags' } },
						{ label: 'アーキテクチャ', slug: 'reference/architecture', translations: { en: 'Architecture' } },
					],
				},
			],
		}),
	],
});
