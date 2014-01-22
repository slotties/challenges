package solutions.coinDenomination;

import java.io.IOException;
import java.io.StringReader;
import java.util.Arrays;
import java.util.Collection;
import java.util.Scanner;

/*
 * http://www.reddit.com/r/dailyprogrammer/comments/1q18a5/110613_challenge_134_intermediate_coin/
 */
public class CLI {
	public static void main(String[] args) throws IOException {
		StringReader data = new StringReader(
				"2 8\n" +
				"1 5 10 25 50 100\n" +
				"1 2 5 10 20 50 100 200\n"
				);
		
		try (Scanner in = new Scanner(data)) {
			int numOfCurrencyDefs = in.nextInt();
			int currencyValue = in.nextInt();
			in.nextLine();
			
			CoinDenominator denominator = new CoinDenominator();
			
			for (int i = 0; i < numOfCurrencyDefs; i++) {
				String line = in.nextLine();
				int[] coinValues = readCoinValues(line);

				System.out.printf("Currency %d Combinations:\n", i + 1);
				Collection<int[]> combinations = denominator.allCombinations(currencyValue, coinValues);
				printCombinations(combinations, coinValues);
			}
		}
	}
	
	private static void printCombinations(Collection<int[]> combinations, int[] coinValues) {
		for (int[] combination : combinations) {
			for (int i = 0; i < combination.length; i++) {
				if (combination[i] > 0) {
					System.out.printf("%d:%d ", coinValues[i], combination[i]);
				}
			}
			
			System.out.println();
		}
	}
	
	private static int[] readCoinValues(String line) {
		String[] coinStrings = line.split("\\s+");
		int[] coinValues = new int[coinStrings.length];
		for (int i = 0; i < coinStrings.length; i++) {
			coinValues[i] = Integer.valueOf(coinStrings[i]);
		}
		
		Arrays.sort(coinValues);
		
		return coinValues;
	}
}
