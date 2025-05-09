//测试Get方法
func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1,", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1 = 123 failed")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key 2 failed")
	}
}

//当使用内存超过了设定值时，是否会触发“无用”节点的移除
//这里改成移除成功的输出，更直观
func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "k3"
	v1, v2, v3 := "value1", "value2", "v3"
	cap := len(k1 + k2 + v1 + v2)
	lru := New(int64(cap), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("key1"); !ok || lru.Len() == 2 {
		t.Fatalf("Removeoldest key1 success!")
	}
}

//回调函数能否被调用
func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	lru := New(int64(10), callback)
	lru.Add("key1", String("123456"))
	lru.Add("k2", String("k2"))
	lru.Add("k3", String("k3"))
	lru.Add("k4", String("k4"))

	expect := []string{"key1", "k2"}

	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expect)
	}
}
