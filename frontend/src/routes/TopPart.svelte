<script>
	import { FontAwesomeIcon } from '@fortawesome/svelte-fontawesome';
	import { icons, colorMapper } from './helpers.js';
	import { clickedTickers } from './store';
	const items = [
		{ icon: icons.EXI, ticker: 'EXI', description: 'Industry' },
		{ icon: icons.IXC, ticker: 'IXC', description: 'Energy' },
		{ icon: icons.IXG, ticker: 'IXG', description: 'Finance' },
		{ icon: icons.IXJ, ticker: 'IXJ', description: 'Health' },
		{ icon: icons.IXN, ticker: 'IXN', description: 'Tech' },
		{ icon: icons.IXP, ticker: 'IXP', description: 'Telecom' },
		{ icon: icons.JXI, ticker: 'JXI', description: 'Utility' },
		{ icon: icons.KXI, ticker: 'KXI', description: 'Staples' },
		{ icon: icons.MXI, ticker: 'MXI', description: 'Materials' },
		{ icon: icons.REET, ticker: 'REET', description: 'RealEst' },
		{ icon: icons.RXI, ticker: 'RXI', description: 'Disc' }
	];

	function toggleClick(ticker) {
		clickedTickers.update((tickers) => {
			const newTickers = new Set(tickers);
			if (newTickers.has(ticker)) {
				newTickers.delete(ticker);
			} else {
				newTickers.add(ticker);
			}
			return newTickers;
		});
	}
</script>

<div class="grid grid-cols-[repeat(auto-fit,minmax(5rem,1fr))] gap-4 m-4">
	{#each items as item}
		<div
			class="flex flex-col items-center justify-center p-4 border border-gray-200 shadow-lg rounded-lg w-20 h-20 cursor-pointer transition hover:bg-gray-100"
			on:click={() => toggleClick(item.ticker)}
		>
			{#if $clickedTickers.has(item.ticker)}
				<div class="text-2xl mb-2" style="color: {colorMapper[item.ticker]};">
					<FontAwesomeIcon icon={item.icon} />
				</div>
			{:else}
				<div class="text-2xl text-gray-400 mb-2">
					<FontAwesomeIcon icon={item.icon} />
				</div>
			{/if}

			<div class="text-sm font-semibold">{item.ticker}</div>
			<div class="text-xs text-gray-600">{item.description}</div>
		</div>
	{/each}
</div>
