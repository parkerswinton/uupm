package textarea

func (m *Model) removePrevious() {
	if len(m.correct) > 0 && m.correct[len(m.correct)-1] == m.cursor {
		m.correct = m.correct[:len(m.correct)-1]
	} else {
		m.incorrect = m.incorrect[:len(m.incorrect)-1]
	}
}

func (m *Model) addCurrent(msg string) {
	if msg == string(m.target[m.cursor]) {
		m.correct = append(m.correct, m.cursor)
	} else {
		m.incorrect = append(m.incorrect, m.cursor)
	}
}

func (m *Model) calcWPM() float32 {
	correct, incorrect, time := float32(len(m.correct)), float32(len(m.incorrect)), float32(m.stopwatch.Elapsed().Minutes())
	return (((correct + incorrect) / 5.0) - incorrect) / time
}
