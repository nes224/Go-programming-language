// Given a string s, find the length of the longest 
// substring
//  without repeating characters.
// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

package main

func lengthOfLongestSubstring(s string) int {
    charIndex := make(map[byte]int)
    maxLength := 0
    start := 0
    for end := 0; end < len(s); end ++ {
        if index, exists := charIndex[s[end]]; exists && index >= start {
            start = index + 1
        }

        charIndex[s[end]] = end
        maxLength = max(maxLength, end - start + 1)
    }
    return maxLength
}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {

}