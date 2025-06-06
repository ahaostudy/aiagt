package summary

import (
	"github.com/aiagt/aiagt/pkg/hash/hset"
	"math"
	"strings"
	"unicode"
)

var (
	separators = hset.FromValues('，', '。', '！', '？', ',', '.', '!', '?', '\n')
)

// Sentence 表示一个句子及其权重
type Sentence struct {
	Text   string
	Weight float64
	Index  int
}

// 分割文本为句子
func splitSentences(text string) []string {
	var sentences []string
	var currentSentence strings.Builder

	for _, char := range text {
		currentSentence.WriteRune(char)
		if separators.Has(char) {
			sentences = append(sentences, strings.TrimSpace(currentSentence.String()))
			currentSentence.Reset()
		}
	}

	if currentSentence.Len() > 0 {
		sentences = append(sentences, strings.TrimSpace(currentSentence.String()))
	}

	return sentences
}

// 预处理句子：去除标点符号并转为小写
func preprocessSentence(sentence string) string {
	var builder strings.Builder
	for _, char := range sentence {
		if unicode.IsLetter(char) || unicode.IsNumber(char) || char == ' ' {
			builder.WriteRune(unicode.ToLower(char))
		}
	}
	return builder.String()
}

// 计算两个句子的余弦相似度
func cosineSimilarity(a, b string) float64 {
	wordsA := strings.Fields(a)
	wordsB := strings.Fields(b)

	// 构建词频向量
	vecA := make(map[string]int)
	vecB := make(map[string]int)

	for _, word := range wordsA {
		vecA[word]++
	}
	for _, word := range wordsB {
		vecB[word]++
	}

	// 计算点积
	dotProduct := 0.0
	for word, countA := range vecA {
		countB := vecB[word]
		dotProduct += float64(countA * countB)
	}

	// 计算向量长度
	magnitudeA := 0.0
	for _, count := range vecA {
		magnitudeA += float64(count * count)
	}
	magnitudeA = math.Sqrt(magnitudeA)

	magnitudeB := 0.0
	for _, count := range vecB {
		magnitudeB += float64(count * count)
	}
	magnitudeB = math.Sqrt(magnitudeB)

	if magnitudeA == 0 || magnitudeB == 0 {
		return 0
	}
	return dotProduct / (magnitudeA * magnitudeB)
}

// TextRank 算法生成文摘
func TextRank(sentences []string, dampingFactor float64, iterations int, topN int) string {
	n := len(sentences)
	similarityMatrix := make([][]float64, n)
	for i := range similarityMatrix {
		similarityMatrix[i] = make([]float64, n)
	}

	// 构建相似度矩阵
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			similarity := cosineSimilarity(
				preprocessSentence(sentences[i]),
				preprocessSentence(sentences[j]),
			)
			similarityMatrix[i][j] = similarity
			similarityMatrix[j][i] = similarity
		}
	}

	// 初始化权重
	weights := make([]float64, n)
	for i := range weights {
		weights[i] = 1.0
	}

	// 迭代更新权重
	for iter := 0; iter < iterations; iter++ {
		newWeights := make([]float64, n)
		for i := 0; i < n; i++ {
			sum := 0.0
			for j := 0; j < n; j++ {
				if i != j && similarityMatrix[i][j] > 0 {
					sum += similarityMatrix[i][j] * weights[j]
				}
			}
			newWeights[i] = (1 - dampingFactor) + dampingFactor*sum
		}
		weights = newWeights
	}

	// 选择权重最高的句子
	rankedSentences := make([]Sentence, n)
	for i, sentence := range sentences {
		rankedSentences[i] = Sentence{Text: sentence, Weight: weights[i], Index: i}
	}

	// 按权重排序
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if rankedSentences[i].Weight < rankedSentences[j].Weight {
				rankedSentences[i], rankedSentences[j] = rankedSentences[j], rankedSentences[i]
			}
		}
	}

	// 选择前 topN 个句子
	var summary []string
	for i := 0; i < topN && i < n; i++ {
		summary = append(summary, rankedSentences[i].Text)
	}

	return strings.Join(summary, " ")
}
