# Weighted Random Hit

The `weightedrandomhit` package is a Go package that provides the ability to determine a 'hit' or 'miss' for a given category based on weighted randomization in a stateless way. The package is particularly useful where the volume of traffic is very high and keeping track of the traffic volume for different types of datapoints can be very expensive. It uses weights assigned to different categories and configuration options to calculate whether a particular category is hit. The package offers options to adjust the chances of hitting a category and to set a default weight for unknown categories.

## Example Scenario

Consider a scenario where you have multiple data streams with different high volumes of datapoints. You want to randomly select samples from each stream to store in a database for validation purposes. However, due to the imbalance in the number of datapoints consumed by each stream, you want to give more chances of selection to streams with fewer datapoints and vice versa. Also, you do not want to keep track of the volume of different incoming streams, which can be very costly.

By using the `weightedrandomhit` package, you can assign weights to each stream category based on their relative datapoint consumption. For example, you can assign a higher weight to a stream with fewer datapoints and a lower weight to a stream with more datapoints. Then, you can specify the target number of samples you want to collect per day (e.g., 5 samples) and the maximum allowed hits for each stream.

The package will perform weighted randomization to determine if a particular stream should be sampled based on its assigned weight and the overall target hit count. This ensures that the sampling is stateless, simple, and doesn't become a performance bottleneck while considering the approximate imbalance in the number of datapoints consumed by each stream.

## How It Works

The `weightedrandomhit` package uses a weighted randomization algorithm to determine if a category should be considered a 'hit' or 'miss'. The algorithm works as follows:

1. For each category, calculate the probability of a hit based on its assigned weight and the total weight of all categories.
2. Generate a random number between the given range.
3. If the random number is less than or equal to the calculated probability for a category, consider it a 'hit'.
4. If the number of hits for a category exceeds the maximum allowed hits, consider it a 'miss'.

The package also provides configuration options to adjust the chances of hitting a category and to set a default weight for unknown categories.

## Installation

To install the `weightedrandomhit` package, use the following command:

```bash
go get github.com/grab/weighted-random-hit
```

## Usage

Here's an example of how to use the `weightedrandomhit` package:

```go
package main

import (
	"fmt"
	weightedrandomizer "github.com/grab/weighted-random-hit"
)

func main() {
	categoryWeights := map[string]int{
		"stream1": 10,
		"stream2": 20,
		"stream3": 30,
	}

	targetHit := 5
	maxAllowedHitsForEachCategory := 10

	category := "stream2"
	isHit := weightedrandomizer.IsCategoryHit(category, categoryWeights, targetHit, maxAllowedHitsForEachCategory)
	fmt.Printf("Category %s hit: %v\n", category, isHit)
}
```

In this example, we define a map of category weights, where each category represents a data stream. The weights are assigned based on the approximate imbalance in the number of datapoints consumed by each stream. We then specify the target hit count (e.g., 5 samples a day) and the maximum allowed hits for each category. Finally, we call the `IsCategoryHit` function with a specific stream category to determine if it should be sampled based on the weighted randomization.

The package also provides additional options to customize the behavior:

- `WithChanceMultiplier(multiplier int)`: Multiplies the chance of a category hit. Useful for testing purposes.
- `WithDefaultWeight(defaultWeight int)`: Sets a default weight for unknown categories.

Here's an example of using these options:

```go
defaultWeight := 5
chanceMultiplier := 2

weightedrandomizer.WithDefaultWeight(defaultWeight)
weightedrandomizer.WithChanceMultiplier(chanceMultiplier)
```

## Configuration

The `weightedrandomhit` package allows you to configure the following options:

- `defaultWeight` (default: 0): The default weight assigned to unknown categories. If not set, unknown categories will be ignored.
- `chanceMultiplier` (default: 1): A multiplier to adjust the chances of hitting a category. Useful for testing purposes.

## Performance Considerations

The `weightedrandomhit` package is designed to be lightweight and efficient. The algorithm has a time complexity of O(1) for determining a hit or miss for a given category. The space complexity is O(n), where n is the number of categories.

However, it's important to note that the package relies on random number generation, which can be a potential performance bottleneck if called very frequently. In such cases, consider using a more optimized random number generator or caching the results if possible.

## License

This package is open-source and available under the [MIT License](LICENSE).

## Credits

The package is authored by [Riyadh Sharif](https://github.com/riyadhctg) and [Jialong Loh](https://github.com/jlloh)