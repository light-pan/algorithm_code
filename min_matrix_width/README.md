```
给定一个矩阵，包含N*M个整效，和一个包含K个整数的数组。
现在要求在这个矩阵中找一个宽度最小的子矩阵，要求子短阵包含数组中所有的整数。
```
### 输入描述：
```
第一行输入两个正整数N，M，表示矩阵大小。
接下来N行M列表示矩阵内容。
下一行包含一个正整数K。
下一行包含K个整数，表示所需包含的数组，K个整数可能存在重复数字所有输入数据小于1000。
```
### 输出描述：
```
输出包含一个整效，表示满足要求子短阵的最小宽度，若找不到，输出-1.
```
### 示例1
#### 输入
```
2 5
1 2 2 3 1
2 3 2 3 2
3
1 2 3
```
#### 输出
```
2
```