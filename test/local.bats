@test "Running kubectl in the root directory should not error" {
  cd /
  kubectl
}
