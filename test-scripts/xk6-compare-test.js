import compare from 'k6/x/compare';

export default function () {
    console.log(`${compare.isGreater(2, 1)}, ${compare.comparison_result}`);
    console.log(`${compare.isGreater(1, 3)}, ${compare.comparison_result}`);

    const state = compare.getInternalState();
    console.log(
        `Active VUs: ${state.activeVUs}, Iteration: ${state.iteration}, VU ID: ${state.vuID}, VU ID from runtime: ${state.vuIDFromRuntime}`
    );
}
