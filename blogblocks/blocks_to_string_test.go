package blogblocks

import "testing"

func TestBlocksToString(t *testing.T) {
	content := `[{"Id":"201909070914390423","ParentId":"20151018210355990473","Type":"Text","Sequence":1,"Attributes":{"Text":"These are some quick \"coppy-and-paste\" snippets that I use all the time."}},{"Id":"201909070915400111","ParentId":"20151018210355990473","Type":"Heading","Sequence":2,"Attributes":{"Level":"2","Text":"PHP. Report All PHP Errors"}},{"Id":"201909070916500752","ParentId":"20151018210355990473","Type":"Code","Sequence":3,"Attributes":{"Language":"php","Code":"error_reporting(E_ALL);\r\nini_set('display_errors', true);\r\nini_set('display_startup_errors', true);"}},{"Id":"201909071006470870","ParentId":"20151018210355990473","Type":"Heading","Sequence":4,"Attributes":{"Level":"2","Text":"Composer. As Lean and Minimal  As Possible"}},{"Id":"201909071008090321","ParentId":"20151018210355990473","Type":"Code","Sequence":5,"Attributes":{"Language":"php","Code":"composer update --prefer-dist -o -vvv --profile"}},{"Id":"201909071009380954","ParentId":"20151018210355990473","Type":"Heading","Sequence":6,"Attributes":{"Level":"2","Text":"PHP. XML to Array"}},{"Id":"201909071011270286","ParentId":"20151018210355990473","Type":"Code","Sequence":8,"Attributes":{"Language":"php","Code":"function xmlToArray($xml) {\r\n    if (is_string($xml)) {\r\n        $sml = simplexml_load_string($xml, 'SimpleXMLElement', LIBXML_NOCDATA);\r\n    }\r\n\r\n    $array = [];\r\n    $array['TagName'] = $xml->getName();\r\n    $array['TagText'] = trim((string) $xml);\r\n    $array['TagAttributes'] = array();\r\n    $array['TagChildren'] = array();\r\n\r\n    foreach ($xml->attributes() as $k => $v) {\r\n        $array['TagAttributes'][$k] = (string) $v;\r\n    }\r\n\r\n    foreach ($xml->children() as $k => $v) {\r\n        $array['TagChildren'][] = xmlToArray($v);\r\n    }\r\n\r\n    return $array;\r\n}"}},{"Id":"201909071009410140","ParentId":"20151018210355990473","Type":"Heading","Sequence":9,"Attributes":{"Level":"2","Text":"Get Content Between Two X/HTML Tags"}},{"Id":"201909071011470632","ParentId":"20151018210355990473","Type":"Code","Sequence":10,"Attributes":{"Language":"php","Code":"function getContentBetweenTags($tagName, $html){\r\n    $pattern = \"#<$tagName(\\\\s+[^>]+)?>(.*?)</$tagName>#s\";\r\n    preg_match($pattern, $html, $matches);\r\n    return $matches[2];\r\n}"}},{"Id":"201909071012050990","ParentId":"20151018210355990473","Type":"Heading","Sequence":11,"Attributes":{"Level":"2","Text":"PHP. Get Current URL"}}]`
	blocks := BlocksToString(content)

	if blocks == "Error parsing content. Please try again later." {
		t.Error("Error parsing content. Please try again later.")
	}

	t.Log(blocks)
}
