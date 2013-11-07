package solutions.coinDenomination;

import java.util.Arrays;
import java.util.Collection;
import java.util.LinkedList;


public class CoinDenominator {	
	public Collection<int[]> allCombinations(int currencyValue, int[] coinValues) {
		Collection<int[]> combinations = new LinkedList<int[]>();
		collectCombinations(currencyValue, coinValues, 0, new int[coinValues.length], combinations);
		
		return combinations;
	}
	
	private void collectCombinations( int currentCurrencyValue, int[] coinValues, int coinValuesOffset, int[] currentCombination, Collection<int[]> combinations) {
		if (currentCurrencyValue == 0) {
			// Yay, the combinations are fine!
			int[] combination = Arrays.copyOf(currentCombination, currentCombination.length);
			combinations.add(combination);
		} else if (coinValuesOffset < coinValues.length) {
			int maxCoins = currentCurrencyValue / coinValues[coinValuesOffset];
			for (int i = 0; i <= maxCoins; i++) {
				currentCombination[coinValuesOffset] = i;
				collectCombinations(currentCurrencyValue - (i * coinValues[coinValuesOffset]), coinValues, coinValuesOffset + 1, currentCombination, combinations);
			}
			// Reset the temporary/working array.
			currentCombination[coinValuesOffset] = 0;
		}
	}
}
