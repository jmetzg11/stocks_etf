<script>
	import { onMount } from 'svelte';
	import { getData } from './helpers.js';
	import TopPart from './TopPart.svelte';
	import Bars from './Bars.svelte';
	let data = [];
	onMount(() => {
		getData();
	});
</script>

<div class="flex flex-col mx-auto max-w-[1200px] h-screen">
	<div class="flex-none">
		<TopPart />
	</div>
	<div class="flex-none mb-5 z-10">
		<Bars />
	</div>
	<div class="flex-none grid grid-cols-1 md:grid-cols-6 gap-4 px-4 mt-5">
		<div class="md:col-span-4 min-h-[300px] z-20">
			<h1>Line Graph</h1>
		</div>
		<div class="md:col-span-2 overflow-y-auto max-h-[250px] mb-3">
			<h1>full table</h1>
		</div>
	</div>
	<div>
		<h1>Info</h1>
	</div>
</div>
