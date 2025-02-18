
# Range-Based Statistics Application

This application performs dynamic range calculations and statistical analysis on a stream of numbers provided through standard input or a text file. It identifies outliers, adjusts ranges, and computes key statistical metrics, such as average, median, and standard deviation.

---

## **Introduction**

The Range-Based Statistics application is designed to process a series of integers dynamically, ensuring robust range calculations and outlier detection. This tool can be used for applications requiring adaptive thresholds and data insights.

---

## **Usage**

### **Run the Application**
Execute the application by providing numbers via standard input:

```bash
echo -e "10\n20\n30\n40\n50" | go run main.go
```

Alternatively, you can redirect input from a file:

```bash
go run main.go < numbers.txt
```

### **Expected Output**
The program outputs the calculated dynamic range (lower and upper bounds) for each input number in real-time:

```plaintext
<lower_bound> <upper_bound>
```

---

## **Key Functions and Modules**

### **Range Calculations**
- **`GetRange`**: Dynamically calculates the lower and upper bounds based on weighted metrics and statistical adjustments.
- **`IsExtreme`**: Detects whether a number qualifies as an outlier based on standard deviations.
- **`RemoveOutlier`**: Excludes outliers from the list of numbers for recalculating ranges.

### **Statistical Functions**
- **`GetAverage`**: Computes the average of a list of integers.
- **`GetMedian`**: Determines the median from a sorted list of integers.
- **`GetStandardDeviation`**: Calculates the standard deviation using variance.
- **`GetVariance`**: Determines variance by summing squared deviations from the mean.



## **Libraries and Dependencies**

The application leverages the Go standard library:
- **`os`**: For file handling.
- **`bufio`**: For buffered I/O operations.
- **`strconv`**: For converting strings to integers.
- **`fmt`**: For formatted input/output.
- **`math`**: For mathematical calculations.

---

## **Performance Considerations**
- Optimized for real-time processing of dynamic data streams.
- Efficient statistical calculations ensure responsiveness with large inputs.
- Robust outlier management maintains integrity in noisy datasets.

---

## **Limitations**
- Processes integer inputs only.
- Requires correctly formatted data streams; non-integer inputs may result in undefined behavior.
- Assumes reasonable bounds for statistical metrics; extreme variability may affect accuracy.

---

## **Author**
Christos Kosmatos

---

## **Example Workflow**

1. **Input Numbers**: A stream of integers is provided via standard input or a file.
2. **Outlier Management**: Outliers are identified and removed dynamically.
3. **Range Adjustment**: New ranges are calculated based on recent inputs.
4. **Output Results**: The lower and upper bounds are printed for each input.

---

## **Future Enhancements**
- Support for non-integer inputs (e.g., floats).
- Configuration options for customizing outlier thresholds and weight distributions.
- Enhanced logging and debugging tools for large-scale data streams.
