package helper

func ReverseDateYYMMDDtoDDMMYY(date string) string {
    var result string
    year := date[:2]
    month := date[2:4]
    day := date[4:]
    result = day+"/"+month+"/"+year
    return result
}