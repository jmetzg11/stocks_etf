import { writable } from 'svelte/store';

export const barDataStore = writable({
	categories: [],
	data: []
});

export const prepReserveData = (reserves) => {
	const categories = reserves.map((d) => d.Stock);
	const data = reserves.map((d) => d.Value);
	barDataStore.set({ categories, data });
};

export const transactionStore = writable({});

export const clickedTickers = writable(
	new Set(['EXI', 'IXC', 'IXG', 'IXN', 'JXI', 'IXP', 'RXI', 'KXI', 'MXI', 'REET', 'IXJ'])
);
