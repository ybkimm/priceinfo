package main

import "github.com/chromedp/chromedp"

const jsParseDanawaPrices = `
;(() => {
	const getElementText = el => 
		el.nodeType === 3 || el.nodeType === 4
		? el.nodeValue.trim()
		: el.nodeType === 1 && el.tagName.toLowerCase() === 'img'
		? el.getAttribute('alt').trim() || ''
		: el.nodeType === 1
		? Array.from(el.childNodes).map(child => getElementText(child)).join('')
		: ''

	const getPrice = el => {
		const innerText = el.innerText.trim()
		if (innerText === '(무료배송)') {
			return 0
		}
		return parseInt(
			innerText.replace(/^\(배송비 ([0-9,]+)원\)/i, '$1')
				.replace(',', '')
		)
	}

	return Array.from(document.querySelectorAll('#defaultMallList .diff_item'))
		.map(item => {
			const mallName = getElementText(
				item.querySelector('.d_mall > .link')
			)

			const price = parseInt(
				item.querySelector('.d_dsc .price .prc_c ')
					.innerText
					.replace(',', '')
			)

			const shippingFee = getPrice(
				item.querySelector('.d_dsc .ship')
			)

			const url = item.querySelector('.d_buy > a')
				.getAttribute('href')

			return {
				mallName,
				price,
				shippingFee,
				url
			}
		})
})()
`

func init() {
	parserMap["prod.danawa.com"] = func(out *[]PriceInfo) chromedp.Action {
		return actions(
			chromedp.WaitVisible("#defaultMallList .diff_item"),
			chromedp.Evaluate(jsParseDanawaPrices, out),
		)
	}
}
