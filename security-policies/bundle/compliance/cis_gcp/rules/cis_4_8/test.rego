package compliance.cis_gcp.rules.cis_4_8

import data.cis_gcp.test_data
import data.compliance.policy.gcp.data_adapter
import data.lib.test
import future.keywords.if

test_violation if {
	eval_fail with input as rule_input({})
	eval_fail with input as rule_input({"shieldedInstanceConfig": {"enableIntegrityMonitoring": false, "enableVtpm": true}})
	eval_fail with input as rule_input({"shieldedInstanceConfig": {"enableIntegrityMonitoring": true, "enableVtpm": false}})
}

test_pass if {
	eval_pass with input as rule_input({"shieldedInstanceConfig": {"enableIntegrityMonitoring": true, "enableVtpm": true}})
}

test_not_evaluated if {
	not_eval with input as test_data.not_eval_resource
}

rule_input(info) := test_data.generate_compute_resource("gcp-compute-instance", info)

eval_fail if {
	test.assert_fail(finding) with data.benchmark_data_adapter as data_adapter
}

eval_pass if {
	test.assert_pass(finding) with data.benchmark_data_adapter as data_adapter
}

not_eval if {
	not finding with data.benchmark_data_adapter as data_adapter
}
