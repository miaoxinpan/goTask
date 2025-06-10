package main

import (
	"fmt"
	"sort"
)

func main() {
	// 测试用例 只出现一次的数字
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber1(nums))
	fmt.Println(singleNumber2(nums))

	// 测试用例 回文数
	palindrome := 121
	fmt.Println(isPalindrome(palindrome))

	// 测试用例 有效的括号
	bracketStr := "({[]})"
	fmt.Println(isValid(bracketStr))

	// 测试用例 最长公共前缀
	commonPrefix := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(commonPrefix))

	// 测试用例 删除有序数组中的重复项
	nums = []int{1, 1, 2, 3, 3, 4, 4, 5, 5, 5, 5, 5, 6, 6, 6, 7, 10, 11, 111, 123, 124, 124, 125, 125}
	fmt.Println(removeDuplicates(nums))

	// 测试用例 加一
	digits := []int{1, 1, 2}
	fmt.Println(plusOne(digits))

	// 测试用例 合并区间
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))

	// 测试用例 两数之和
	nums = []int{1, 9, 5, 6, 8, 2, 3}
	target := 4
	//暴力解法
	fmt.Println(twoSum(nums, target))
	//利用hash表
	fmt.Println(twoSum2(nums, target))
}

/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*
《只出现一次的数字》
  给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
  你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间

由于每个元素均出现两次  那么就用map  第一次放入 第二次移除的思路 这样map里面就只剩下只出现一次的数字


也可以用最快的方式 异或运算
*/
func singleNumber1(nums []int) int {
	var value int
	maps := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, exists := maps[nums[i]]; exists {
			delete(maps, nums[i]) // 无返回值，key不存在时不会报错
		} else if !exists {
			maps[nums[i]] = 1
		}

	}
	fmt.Println(maps)
	for k, _ := range maps {
		value = k
	}
	return value
}

/*
初始值 res = 0
res = 0 ^ 4 = 4
res = 4 ^ 1 = 5
res = 5 ^ 2 = 7
res = 7 ^ 1 = 6  (因为 0111 ^ 0001 = 0110)
res = 6 ^ 2 = 4  (因为 0110 ^ 0010 = 0100)
最终返回4
*/
func singleNumber2(nums []int) int {
	var value int
	for _, num := range nums {
		value ^= num
	}
	return value
}

/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*
《回文数》
	给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
	回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
	例如，121 是回文，而 123 不是。

	双指针  我读一个  末尾读一个  如果相等 继续向中间走 直到相遇
*/

func isPalindrome(palindrome int) bool {
	//先把x转换成字符串
	var palindromeStr string = fmt.Sprintf("%d", palindrome)
	//双指针，第一个指针从左向右，第二个指针从右向左
	left, right := 0, len(palindromeStr)-1
	//如果左边一直小于右边  那么说明还没相遇
	for left < right {
		//如果不等于 直接返回false
		//如果相等 左边加1 右边减1
		if palindromeStr[left] != palindromeStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*
《有效的括号》
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
	有效字符串需满足：
	1.左括号必须用相同类型的右括号闭合。
	2.左括号必须以正确的顺序闭合。
	3.每个右括号都有一个对应的相同类型的左括号。

	使用栈的特性 先进后出 来解决这个问题
	如果遇到左括号就入栈  如果遇到右括号就出栈  如果出栈的不是对应的左括号 那么就返回false

	但是go没有栈  所以用切片
	这是最常用的方式，利用切片的动态扩容特性：
		入栈：append操作
		出栈：截取切片[:len(slice)-1]
*/

func isValid(bracketStr string) bool {
	//定义一个切片来模拟栈的功能
	stack := []rune{}
	//将所有的括号都存入map中    因为遇到左括号 就得去栈里面看有没有对应的右括号，所以map里面key为左括号  value为右括号
	//都是切片类型  因为直接对这个切片循环 所以省去了类型转换的麻烦
	mapping := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, char := range bracketStr {
		// 左括号全部按顺序入栈
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else if char == ')' || char == '}' || char == ']' {
			// 栈是空的  或者 右括号不匹配，返回 false
			if len(stack) == 0 || stack[len(stack)-1] != mapping[char] {
				return false
			}
			// 右括号匹配，出栈
			stack = stack[:len(stack)-1]
		}
	}
	//循环结束  如果栈为空，则为true
	return len(stack) == 0
}

/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*
《最长公共前缀》
	编写一个函数来查找字符串数组中的最长公共前缀。
	如果不存在公共前缀，返回空字符串 ""。


	思路：先将第一个字符串作为公共前缀，然后遍历后面的字符串，逐个比较公共前缀和当前字符串的前缀，
	如果不匹配，就将公共前缀缩短，直到找到最长的公共前缀为止。


(
func longestCommonPrefix(commonPrefix []string) string {
	//先定义一个空字符串返回
	result := ""
	// 如果传入的字符串数组为空，直接返回空字符串
	if len(commonPrefix) == 0 {
		return result
	}
	// 以第一个字符串作为初始的公共前缀
	result = commonPrefix[0]
	// 遍历后面的字符串
	for i := 1; i < len(commonPrefix); i++ {
		// 比较当前字符串和公共前缀
		for j := 0; j < len(result) && j < len(commonPrefix[i]); j++ {
			// 如果当前字符不匹配，缩短公共前缀
			if result[j] != commonPrefix[i][j] {
				result = result[:j]
				break
			}
		}
	}
	return result
}
)----------------上面这个方法没有通过leetcode测试---------------
下面是ai优化后的代码   学习一下
*/
func longestCommonPrefix(commonPrefix []string) string {
	//如果传入的字符串数组为空，直接返回空字符串
	if len(commonPrefix) == 0 {
		return ""
	}
	//以第一个字符串作为初始的公共前缀
	prefix := commonPrefix[0]
	//循环字符串数组  从第二个字符串开始
	for i := 1; i < len(commonPrefix); i++ {
		// 条件1 ：如果前缀的长度大于零
		// str[i][:?] ?=len(prefix) 例子:flower[:3]="flo" 含头不含尾
		// 条件2 ：当前字符串的长度小于前缀的长度 或者 以 前缀长度为截取字符串的长度,截取出来后跟前缀相比较

		//满足条件1且条件2  就继续for循环 将前缀缩短  一直到不满足这个for循环的条件
		// 假设len(prefix) 是10 len(commonPrefix[i])是3 那就是直接先把prefix缩短到3 然后就是看commonPrefix[i]的前3个字符是否和prefix相等
		// 如果不相等 就将prefix缩短一个字符 然后再看commonPrefix[i]的前2个字符是否和prefix相等
		for len(prefix) > 0 && (len(commonPrefix[i]) < len(prefix) || commonPrefix[i][:len(prefix)] != prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if prefix == "" {
			return ""
		}
	}
	return prefix
}

/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*
《删除有序数组中的重复项》

	给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
	考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
	更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
	返回 k 。
示例1
	输入：nums = [1,1,2]
	输出：2, nums = [1,2,_]
	解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例2
	输入：nums = [0,0,1,1,1,2,2,3,3,4]
	输出：5, nums = [0,1,2,3,4]
	解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素

思路：用两个指针对当前这个数组遍历  由于是按照顺序来的  所以只要当前指针的值和前一个指针的值不相等 就将当前指针的值放到前一个指针的下一个位置 而且只需要返回元素个数，不过我们可以打印一下
*/
func removeDuplicates(nums []int) int {
	//先定义一前一后两个指针 for循环中的i 就是另外一个指针了
	left := 0
	//因为第一个肯定是最小的  所以直接从下标为1开始
	for i := 1; i < len(nums); i++ {
		if !(nums[i-1] == nums[i]) { //如果不相等 让指针走一步
			left++
			//并且还得换数组中数据的位置
			nums[left] = nums[i]
		}
	}
	//因为返回的是元素中的个数 所以下标要加1
	return left + 1
}

/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*
《加一》
	给定一个由 整数 组成的 非空 数组所表示的非负整数，在该数的基础上加一。
	最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
	你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1：
	输入：digits = [1,2,3]
	输出：[1,2,4]
	解释：输入数组表示数字 123。
示例 2：
	输入：digits = [4,3,2,1]
	输出：[4,3,2,2]
	解释：输入数组表示数字 4321。
示例 3：
	输入：digits = [9]
	输出：[1,0]
	解释：输入数组表示数字 9。
	加 1 得到了 9 + 1 = 10。
	因此，结果应该是 [1,0]。

	传入的一个数字 转成数组  比如说 一个整数  123  加 1  再放到数组里面  124
	99 变成 100 放到数组里面

那就是直接从最后面开始  如果是9的 就往前一位加1 并且把当前位变成0
*/
func plusOne(digits []int) []int {
	for i := len(digits); i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	//如果走到这，就代表数组里面的元素都是9 那么最前面就是1 后面都是0  就类似  99  那么就新建一个1的切片  把数组里面00拼接再这个切片后面  变成100
	return append([]int{1}, digits...)
}

/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*
《合并区间》
	以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
示例 1：
	输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
	输出：[[1,6],[8,10],[15,18]]
	解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：
	输入：intervals = [[1,4],[4,5]]
	输出：[[1,5]]
	解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。

	相当于就是看有没有交叉的数字
	如果前一个末尾大于等于后一个的开始  那么就是有重叠  否则就是整个放入新的里面

	还有这种试例  [[1,4],[0,4]]  那就是不能保证第一个就是最小的那个范围  得先排序

	哎呀  还有这种试例  [[1,4],[2,3]]  那就是直接取后面的那个值了  得比一下了
*/
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	// 先按区间起点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//默认第一个区间不重复，所以从下标1开始
	result := [][]int{intervals[0]}
	//记录result的下标
	flag := 0
	for i := 1; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		pre := result[flag][1]
		if pre >= start {
			if pre > end {
				end = pre
			}
			result[flag][1] = end
		} else {
			flag++
			result = append(result, intervals[i])
		}
	}

	return result
}

/*--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------*/
/*
《两数之和》
	给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
	你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
	你可以按任意顺序返回答案。

示例 1：
	输入：nums = [2,7,11,15], target = 9
	输出：[0,1]
	解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：
	输入：nums = [3,2,4], target = 6
	输出：[1,2]
示例 3：
	输入：nums = [3,3], target = 6
	输出：[0,1]

	暴力解法
*/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// nums = []int{1, 7, 5, 6, 8, 2, 3}  map里面放的主要是key是数组中的值  value是值所对应的下标
// map[value]index
// 比如说 target 是 10  他进来先看  hashTable表里面有没有 10-1=9 		没有的话 就先把 key为1  value 为0 存到map里面
// 然后下一个循环 index为1  value为7  这个时候就看hashTable表里面有没有 10-7=3 没有的话 就把key为7 value为1 存到map里面
// 然后又进入下一个循环 一直到 8 进来 index为4  value为8 这个时候就看hashTable表里面有没有 10-8=2 没有的话 就把key为8 value为4 存到map里面
// 然后2  index为5 value为2  这个时候就看hashTable表里面有没有 10-2=8  这个时候  hashTable[8]的值为4  所以 直接return p为4 当前的index为5 所以是 [4,5]
func twoSum2(nums []int, target int) []int {
	//新建一个map用来存可能的答案
	hashTable := map[int]int{}
	for index, value := range nums {
		//如果hashTable 存在  目标值-当前值 直接返回 下标
		if p, ok := hashTable[target-value]; ok {
			return []int{p, index}
		}
		hashTable[value] = index
	}
	return nil
}
