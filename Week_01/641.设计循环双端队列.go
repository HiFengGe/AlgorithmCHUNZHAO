/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 */

// @lc code=start
type MyCircularDeque struct {
	maxLength int   //设置长度
	nums      []int //值
	length    int   //当前长度
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		maxLength: k,
		nums:      make([]int, 0, 10),
		length:    0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	//是否在合理边界中
	if this.length < this.maxLength {
		this.nums = append([]int{value}, this.nums...)
		// 长度+1
		this.length++
		return true
	}
	return false

}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	//是否在合理边界中
	if this.length < this.maxLength {
		this.nums = append(this.nums, value)
		// 长度+1
		this.length++
		return true
	}
	return false

}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.length > 0 {
		this.nums = this.nums[1:]
		this.length--
		return true
	}
	return false
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.length > 0 {
		this.nums = this.nums[:this.length-1]
		this.length--
		return true
	}
	return false

}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.length > 0 {
		return this.nums[0]
	}
	return -1
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.length > 0 {
		return this.nums[this.length-1]
	}
	return -1
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.length == 0 {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.length == this.maxLength {
		return true
	}
	return false
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

