package query

import "testing"

func TestSimpleQuery(t *testing.T) {
	var q Query
	expected := "SELECT t.id,t.name FROM test t WHERE t.name = test;"

	w := Where{
		Field: "t.name",
		Exp:   "=",
		Val:   "test",
	}

	sql := q.Select("test", "t", "t.id", "t.name").Where(w).Build()

	if sql != expected {
		t.Errorf("Wrong sql, given %s, exptected %s", sql, expected)
		t.Fail()
	}

}

func TestComplexQuery(t *testing.T) {
	var q Query

	partnerWhere := Where{
		Field: "pp.partner_id",
		Exp:   "=",
		Val:   "$1",
	}

	countryWhere := Where{
		Field: "pl.default_country_code",
		Exp:   "IN",
		Val:   "($2, \"*\")",
	}

	expected := "SELECT p.id,p.name,p.period,p.is_auto_renew,p.status,p.trial_period_len,p.creation_time,p.membership_type,p.base_package_id,pp.partner_id,pl.default_country_code,pl.localized_name,pl.price_currency_code,pl.price_value,pl.price_decimal_point FROM packages p LEFT JOIN packages_partners pp ON p.id = pp.package_id LEFT JOIN package_localization pl ON p.id = pl.package_id WHERE pp.partner_id = $1 AND pl.default_country_code IN ($2, \"*\");"

	sql := q.Select("packages", "p", "p.id", "p.name", "p.period", "p.is_auto_renew", "p.status", "p.trial_period_len", "p.creation_time",
		"p.membership_type", "p.base_package_id", "pp.partner_id", "pl.default_country_code", "pl.localized_name", "pl.price_currency_code", "pl.price_value", "pl.price_decimal_point",
	).LeftJoin("packages_partners", "pp", "p.id", "pp.package_id",
	).LeftJoin("package_localization", "pl", "p.id",  "pl.package_id",
	).Where(partnerWhere).AndWhere(countryWhere).Build()

	if sql != expected {
		t.Errorf("Wrong sql, given %s, exptected %s", sql, expected)
		t.Fail()
	}
}
