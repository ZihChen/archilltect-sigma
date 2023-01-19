package gpt3service

type Model string

const (
	DavinciModel Model = "text-davinci-003"
	AdaModel     Model = "text-ada-001"
	BabbageModel Model = "text-babbage-001"
	CurieModel   Model = "text-curie-001"
)
