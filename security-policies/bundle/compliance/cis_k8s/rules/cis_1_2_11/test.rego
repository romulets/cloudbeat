package compliance.cis_k8s.rules.cis_1_2_11

import data.compliance.cis_k8s.data_adapter
import data.kubernetes_common.test_data
import data.lib.test
import future.keywords.if

test_violation if {
	eval_fail with input as rule_input("--enable-admission-plugins=AlwaysAdmit")
	eval_fail with input as rule_input("--enable-admission-plugins=PodNodeSelector,AlwaysAdmit")
}

test_pass if {
	eval_pass with input as rule_input("--enable-admission-plugins=RBAC")
	eval_pass with input as rule_input("")
}

test_not_evaluated if {
	not_eval with input as test_data.process_input("some_process", [])
}

rule_input(argument) := test_data.process_input("kube-apiserver", [argument])

eval_fail if {
	test.assert_fail(finding) with data.benchmark_data_adapter as data_adapter
}

eval_pass if {
	test.assert_pass(finding) with data.benchmark_data_adapter as data_adapter
}

not_eval if {
	not finding with data.benchmark_data_adapter as data_adapter
}
