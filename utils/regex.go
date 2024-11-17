package utils

import "regexp"


var sqliErrorRegex = regexp.MustCompile(`(SQL syntax)|(mysql_query)|(pg_query)`) // Example

// ... Add more regex patterns for different vulnerabilities
