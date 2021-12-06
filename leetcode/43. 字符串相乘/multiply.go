package main

import "fmt"

/*
class Solution {
public:
    void addIndex(string& str, int index, int val) {
        int sum = str[index] - '0' + val;
        str[index] = sum % 10 + '0';
        if (sum > 9) {
            addIndex(str, index - 1, sum/10);
        }
    }

    string multiply(const string& num1, const string& num2) {
        if (num1 == "0" || num2 == "0") {
            return "0";
        }

        string ret(num1.size() + num2.size(), '0');
        for (int i = num1.size() - 1; i >= 0; --i) {
            for (int j = num2.size() - 1; j >= 0; --j) {
                int product = (num1[i] - '0') * (num2[j] - '0');
                addIndex(ret, i + j + 1, product);
            }
        }

        if (ret[0] == '0')
            ret.erase(ret.begin());

        return ret;
    }
};
*/

func addIndex(str *[]byte, index int, val int) {
	sum := int((*str)[index]) + val
	(*str)[index] = uint8(sum % 10)
	if sum > 9 {
		addIndex(str, index - 1, sum / 10)
	}
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	ret := make([]byte, len(num1) + len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			product := int(num1[i] - '0') * int(num2[j] - '0')
			addIndex(&ret, i + j + 1, product)
		}
	}

	if ret[0] == byte(0) {
		ret = ret[1:]
	}

	for i := 0; i < len(ret); i++ {
		ret[i] = ret[i] + '0'
	}

	return string(ret)
}

func main() {
	ret := multiply("123", "123")
	fmt.Println(ret)
}