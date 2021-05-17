package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	sample_json := `{"filter":{"context":null,"query":"vapormax","original_query":"vapormax","sizes":[],"price_ranges":[],"colors":[],"brands":[],"sort":"relevancy","shops":[],"country_exclusions":[],"categories":[],"division":null,"department":null,"class":null,"subclass":null,"include_persistent":true,"include_flash":false,"sku":null,"upc":null,"anyUpc":null,"page":1,"limit":30,"navItemPermanentId":[],"experiment":null,"style_id_colors":[],"site":null,"nestedColors":false,"groupSkus":false,"includeStoreInventoryInfo":false},"search_cluster":{"name":"fusion","doc_ids":[],"filter_pairs":["query/vapormax","original_query/vapormax","sort/relevancy","include_persistent/true","include_flash/false","page/1","limit/30","nestedColors/false","groupSkus/false","includeStoreInventoryInfo/false"],"host":"http://prod2-fusion","num_found":1,"query_id":"FDWpUMuUOZ","start":0},"page":1,"limit":30,"pages":1,"total":1,"corrected_query":null,"_links":{"self":{"href":"/api/search2/catalog/search?query=vapormax&sort=relevancy&page=1&limit=30&includeFlash=false&includePersistent=true&nestedColors=false"},"first":{"href":"/api/search2/catalog/search?query=vapormax&sort=relevancy&page=1&limit=30&includeFlash=false&includePersistent=true&nestedColors=false"},"last":{"href":"/api/search2/catalog/search?query=vapormax&sort=relevancy&page=1&limit=30&includeFlash=false&includePersistent=true&nestedColors=false"},"http://hautelook.com/rels/search":{"templated":true,"href":"/api/search2/catalog/search{?query,context,division,department,class,subclass,page,includePersistent,includeFlash,limit,sort,sizes%5B%5D*,colors%5B%5D*,priceRanges%5B%5D*,brands%5B%5D*}"},"http://hautelook.com/rels/price-ranges":{"href":"/api/search2/catalog/search/price-ranges-data?query=vapormax&sort=relevancy&page=1&limit=30&includeFlash=false&includePersistent=true&nestedColors=false"},"http://hautelook.com/rels/colors":{"href":"/api/search2/catalog/search/colors-data?query=vapormax&sort=relevancy&page=1&limit=30&includeFlash=false&includePersistent=true&nestedColors=false"},"http://hautelook.com/rels/sizes":{"href":"/api/search2/catalog/search/sizes-data?query=vapormax&sort=relevancy&page=1&limit=30&includeFlash=false&includePersistent=true&nestedColors=false"},"http://hautelook.com/rels/brands":{"href":"/api/search2/catalog/search/brands-data?query=vapormax&sort=relevancy&page=1&limit=30&includeFlash=false&includePersistent=true&nestedColors=false"},"http://hautelook.com/rels/sorts":{"href":"/api/search2/catalog/search/sort-data"},"http://hautelook.com/rels/classification":{"href":"/api/search2/catalog/classification?query=vapormax&nestedColors=false"}},"_embedded":{"http://hautelook.com/rels/products":[{"division":"Kids","department":"Boys' Shoes","class":"Sneakers & Athletic","sub_class":"Athletic","brand_name":"Nike","name":"Air Vapormax 2019","style_num":"AJ2616-002","style_id":3575073,"web_style_id":null,"source":"offer","style_id_color":"3575073-purepltinm/blk-wht-team_orange","fiber_content":"","material":"textile/manmade upper and lining, manmade sole","max_items_per_person":0,"weight":1,"additional_details":"","additional_information":"","more_information":"","description":"With a signature bubbly sole this Vapormax sneaker will give them a sporty, modern look.\n\nY=youth sizing\n\n- Round toe\n- Lace-up vamp\n- Back pull-tab\n- Side logo detail\n- Air pocket sole\n- Imported","care":"","return":"non_returnable","clearance":false,"choke_hazard":false,"qualifies_free_shipping":true,"ship_code":"owned_inventory","ship_time":"3 - 5 days","ship_description":"This item is ready to ship.","shops":[],"country_exclusions":["AU","CA"],"secondary_classification_ids":[],"is_price_visible":false,"_links":{"alternate":{"templated":true,"href":"//www.nordstromrack.com/s/n3575073{?color,size}"},"self":{"href":"/api/nrhl/products/3575073?color=PUREPLTINM%2FBLK-WHT-TEAM+ORANGE"},"http://hautelook.com/rels/skus":{"href":"/api/nrhl/products/3575073/skus"},"http://hautelook.com/rels/size-chart":{"href":"https://fastly.hautelookcdn.com/assets/sizechart/sizechart-kidsshoes_2019.png"},"http://hautelook.com/rels/measurement-guide":{"href":"https://fastly.hautelookcdn.com/assets/measurementguide/"},"offer":{"href":"https://prod-nrhl.offer.vip.nordstrom.com/offer/3575073","type":"NRHL"},"metadata":{"href":"/api/search3/catalog/products/3575073?is_nrhl_style_id=1"}},"_embedded":{"http://hautelook.com/rels/skus":[{"sku":22057634,"upc":"439115733662","nupc":"439115733662","price_retail":170.0,"price_sale":107.97,"regular_price":107.97,"in_clear_the_rack":false,"price_discount":0.36488235,"size":"5Y","standard_size":"5 (Toddler or Big Kid)","color":"PUREPLTINM/BLK-WHT-TEAM ORANGE","color_family":"Black","inventory_level":"high","low_quantity":8,"is_returnable":true,"is_returnable_to_rack":true,"returnability":"returnable","is_clearance":false,"is_final_sale":false,"ship_code":"owned_inventory","ship_time":"3 - 5 days","ship_description":"This item is ready to ship.","_links":{"self":{"href":"/api/nrhl/products/skus/22057634"},"http://hautelook.com/rels/product":{"href":"/api/nrhl/products/3575073"},"http://hautelook.com/rels/swatch":{"href":"https://fastly.hautelookcdn.com/products/AJ2616-002/purepltinm/blk-wht-team_orange.jpg"},"http://hautelook.com/rels/images":[{"href":"https://n.nordstrommedia.com/id/sr3/6403c8f6-e297-42dc-88cd-e7f19c395762.jpeg?width={width}&height={height}","templated":true},{"href":"https://n.nordstrommedia.com/id/sr3/ec113954-a773-41ad-9f10-35be49ce6058.jpeg?width={width}&height={height}","templated":true},{"href":"https://n.nordstrommedia.com/id/sr3/f680e430-f8b5-493b-b074-39a80e4bb401.jpeg?width={width}&height={height}","templated":true},{"href":"https://n.nordstrommedia.com/id/sr3/c88b51d7-0c3e-4852-a5e9-4bf7f648c498.jpeg?width={width}&height={height}","templated":true},{"href":"https://n.nordstrommedia.com/id/sr3/aa469d22-2b81-4765-9630-61282a90ef99.jpeg?width={width}&height={height}","templated":true}],"http://hautelook.com/rels/original-images":[{"href":"https://n.nordstrommedia.com/id/sr3/6403c8f6-e297-42dc-88cd-e7f19c395762.jpeg","templated":true},{"href":"https://n.nordstrommedia.com/id/sr3/ec113954-a773-41ad-9f10-35be49ce6058.jpeg","templated":true},{"href":"https://n.nordstrommedia.com/id/sr3/f680e430-f8b5-493b-b074-39a80e4bb401.jpeg","templated":true},{"href":"https://n.nordstrommedia.com/id/sr3/c88b51d7-0c3e-4852-a5e9-4bf7f648c498.jpeg","templated":true},{"href":"https://n.nordstrommedia.com/id/sr3/aa469d22-2b81-4765-9630-61282a90ef99.jpeg","templated":true}]},"is_price_visible":false,"full_line_transfer":true,"epm_sku_id":"2061614155"}]}}],"http://hautelook.com/rels/price-ranges":{"price_ranges":[{"value":"100-200","label":"$100-$200","count":1,"price_range":{"name":"100-200","from":100.0,"to":200.0}}],"_links":{"self":{"href":"/api/search2/catalog/search/price-ranges-data?query=vapormax&sort=relevancy&page=1&limit=30&includeFlash=false&includePersistent=true&nestedColors=false"}}},"http://hautelook.com/rels/sizes":{"sizes":[{"value":"5 (Toddler or Big Kid)","label":"5 (Toddler or Big Kid)","count":1},{"value":"5.5 (Toddler or Big Kid)","label":"5.5 (Toddler or Big Kid)","count":1}],"brands":[{"value":"5 (Toddler or Big Kid)","label":"5 (Toddler or Big Kid)","count":1},{"value":"5.5 (Toddler or Big Kid)","label":"5.5 (Toddler or Big Kid)","count":1}],"_links":{"self":{"href":"/api/search2/catalog/search/sizes-data?query=vapormax&sort=relevancy&page=1&limit=30&includeFlash=false&includePersistent=true&nestedColors=false"}}},"http://hautelook.com/rels/colors":{"colors":[{"value":"Black","label":"Black","count":1}],"_links":{"self":{"href":"/api/search2/catalog/search/colors-data?query=vapormax&sort=relevancy&page=1&limit=30&includeFlash=false&includePersistent=true&nestedColors=false"}}},"http://hautelook.com/rels/brands":{"brands":[{"value":"Nike","label":"Nike","count":1}],"_links":{"self":{"href":"/api/search2/catalog/search/brands-data?query=vapormax&sort=relevancy&page=1&limit=30&includeFlash=false&includePersistent=true&nestedColors=false"}}},"http://hautelook.com/rels/sorts":{"sorts":[{"value":"relevancy","name":"Relevance","label":"Relevance"},{"value":"featured","name":"Newest Arrivals","label":"Newest Arrivals"},{"value":"most_popular","name":"Best Sellers","label":"Best Sellers"},{"value":"best_value","name":"Best Value","label":"Best Value"},{"value":"price_asc","name":"Price Low To High","label":"Price Low To High"},{"value":"price_desc","name":"Price High To Low","label":"Price High To Low"}],"_links":{"self":{"href":"/api/search2/catalog/search/sort-data"}}}}}`

	toBytes := []byte(sample_json)

	filter1 := filter{}

	err := json.Unmarshal(toBytes, &filter1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(filter1.Filter.Query)

}

type filter struct {
	Filter struct {
		Query string `json:"query"`
	}
}

func GetRequest(requestURL string) (string, error) {
	resp, err := http.Get(requestURL)

	if err != nil {
		fmt.Println("Error has occured")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return string(body), err
}
