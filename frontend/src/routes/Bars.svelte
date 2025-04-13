<script>
	import { onMount } from 'svelte';
	import { FontAwesomeIcon } from '@fortawesome/svelte-fontawesome';
	import { barDataStore, clickedTickers } from './store';
	import { icons, colorMapper } from './helpers';
	let chartId = 'etf-bar-graph';

	const barData = $derived($barDataStore);
	const activeCategories = $derived($clickedTickers);
	const activeCategoriesArray = $derived([...activeCategories]);
	const categories = $derived(barData.categories);
	const data = $derived(barData.data);

	let chart;
	let filteredBarData = $state({
		categories: [],
		data: []
	});
	let barColors = $state([]);
	let barPositions = $state([]);

	$effect(() => {
		filteredBarData.categories = barData.categories.filter((category) =>
			activeCategories.has(category)
		);
		filteredBarData.data = barData.data.filter((_, index) =>
			activeCategories.has(barData.categories[index])
		);
	});
	$effect(() => {
		barColors.length = 0;
		filteredBarData.categories.forEach((c) => barColors.push(colorMapper[c]));
	});

	$effect(() => {
		$inspect(filteredBarData.data);
		$inspect(filteredBarData.categories);
		$inspect(barColors);
	});

	$effect(() => {
		if (chart && filteredBarData.data.length > 0) {
			updateChart();
		}
	});

	onMount(() => {
		console.log('starting bar chart');
		// initializeChart();
	});
</script>

<div class="relative h-full w-full">
	<div
		class="absolute top-0 left-0 w-full text-center font-semibold text-xl md:text-2xl lg:text-3xl xl:text-4xl"
	>
		ETF Reserves
	</div>
	<div id={chartId}></div>
	<div class="absolute bottom-0 left-0 w-full flex justify-around">
		{#each filteredBarData.categories as category, i (category)}
			{#if barPositions[i]}
				<div
					key={category}
					class="absolute"
					style="transform: translateX(-50%); top: 100%; left: {barPositions[
						i
					]}px; color: {colorMapper[category]};"
				>
					<FontAwesomeIcon id={category} key={category} icon={icons[category]} class="text-2xl" />
				</div>{/if}
		{/each}
	</div>
</div>
