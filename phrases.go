package main

import (
	"math/rand"
)

var badWords = []string{
	"в очко себе это надиктуй",
	"что за ебанный писк?",
	"у меня кошка, когда блюет, ито звучит лучше",
	"судя по интонации у тебя явные признаки ДЦП",
	"В рот тебе ноги за твои голосовухи, мразота",
	"Закончило со своей речью, мудило? Умничка, теперь топай нахуй. Топ-топ, блять. Топ-топ, нахуй",
	"А может люди не имеют возможности тебя слушать, есть такой вариант? Но тебе же похуй, так и пизди с собой дальше",
	"Ооооооооо, моя оборооооона! Ребята, послушайте лучше музыку, чем высеры в этом бессмысленном ебучем голосовом сообщении от нашего дорогого друга",
	"Блять, ну какого хуя? Какого хуйца ты это делаешь? Займи свой рот чем-нибудь интересным и пиши буковками",
	"Ублюдок, мать твою, а ну, иди сюда, говно собачье! Что, решил к нам лезть со своими голосовыми?! Ты, засранец вонючий, мать твою! А ну, иди сюда, попробуй меня трахнуть! Я тебя сам трахну, ублюдок! Онанист чёртов, будь ты проклят! Иди, идиот, трахать тебя и всю твою семью! Говно собачье, жлоб вонючий, дерьмо, сука, падла! Иди сюда, мерзавец, негодяй, гад! Иди сюда, ты, говно, жопа!",
	"Извините, что вмешиваюсь, но, господа, не могли бы вы перестать отправлять голосовые сообщения, ну в самом деле. Созвонитесь, пожалуйста, друг с другом и пообщайтесь. Благодарю за внимание, приятного дня.",
}

// GetBadWord ...
func GetBadWord() string {
	return badWords[rand.Intn(len(badWords))]
}
