package mode

// 定义一个结构体
type student struct {
	Name  string
	score float64
}

// 因为student结构体首字母是小写,因此是只能在model使用
// 通过工厂模式来解决

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s,
	}
}

// score字段小写
func (s *student) GetScore() float64 {
	s.score = 5
	return s.score
}
