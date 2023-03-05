package defaults

func Labels() []string {
	return []string{
		"kubernetes.io/arch",
		"kubernetes.io/os",
		"node.kubernetes.io/instance-type",
		"node_pool",
	}
}
