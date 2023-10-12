import Spinner from './loadingSpinner.svelte';

export const loader = (node: any, loading: { subscribe: (arg0: (loading: any) => void) => any }) => {
	let Spin: Spinner<any, any, any> | undefined, loadingValue;
	const unsubscribe = loading.subscribe((loading: any) => {
		if (loading) {
			Spin = new Spinner({
				target: node,
				intro: true
			});
		} else {
			if (Spin) {
				Spin?.$destroy?.();
				Spin = undefined;
			}
		}
	});
};