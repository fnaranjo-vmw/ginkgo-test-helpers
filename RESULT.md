```
Running Suite: Helpers - /ginkgo-test-helper
===================================================================================
Random Seed: 1686057348

Will run 5 of 5 specs
------------------------------
• [FAILED] [0.000 seconds]
Using helper functions in tests Ginkgo when assertions are used in the helper function [It] prints USELESS line number in colorized text and summaries
/ginkgo-test-helper/main_test.go:11

  [FAILED] Something is failing
  In [It] at: /ginkgo-test-helper/helpers.go:11 @ 06/06/23 15:15:48.594
------------------------------
• [FAILED] [0.000 seconds]
Using helper functions in tests Ginkgo when assertions are used in the test body [It] prints USEFUL line number in colorized text and summaries
/ginkgo-test-helper/main_test.go:16

  [FAILED] Something is failing
  In [It] at: /ginkgo-test-helper/main_test.go:19 @ 06/06/23 15:15:48.594
------------------------------
• [FAILED] [0.000 seconds]
Using helper functions in tests Gomega when assertions are used in the helper function [It] prints USELESS line number but USEFUL automatic text in colorized text and summaries
/ginkgo-test-helper/main_test.go:26

  [FAILED] Expected
      <string>: some-value
  to equal
      <string>: never ever fail
  In [It] at: /ginkgo-test-helper/helpers.go:19 @ 06/06/23 15:15:48.594
------------------------------
• [FAILED] [0.000 seconds]
Using helper functions in tests Gomega when assertions are used in the test body [It] prints USEFUL line number but MISSES automatic text in colorized text and summaries
/ginkgo-test-helper/main_test.go:31

  [FAILED] Unexpected error:
      <*errors.errorString | 0xc000304870>: 
      expecting "never ever fail" but got "some-value"
      {
          s: "expecting \"never ever fail\" but got \"some-value\"",
      }
  occurred
  In [It] at: /ginkgo-test-helper/main_test.go:33 @ 06/06/23 15:15:48.595
------------------------------
• [FAILED] [0.000 seconds]
Using helper functions in tests Gomega when the helper function returns an GomegaMatcher [It] prints USEFUL line number and USEFUL automatic text in colorized text and summaries. Getting the best of both worlds
/ginkgo-test-helper/main_test.go:37

  [FAILED] Expected
      <string>: some-value
  to equal
      <string>: never ever fail
  In [It] at: /ginkgo-test-helper/main_test.go:39 @ 06/06/23 15:15:48.595
------------------------------

Summarizing 5 Failures:
  [FAIL] Using helper functions in tests Ginkgo when assertions are used in the helper function [It] prints USELESS line number in colorized text and summaries
  /ginkgo-test-helper/helpers.go:11
  [FAIL] Using helper functions in tests Ginkgo when assertions are used in the test body [It] prints USEFUL line number in colorized text and summaries
  /ginkgo-test-helper/main_test.go:19
  [FAIL] Using helper functions in tests Gomega when assertions are used in the helper function [It] prints USELESS line number but USEFUL automatic text in colorized text and summaries
  /ginkgo-test-helper/helpers.go:19
  [FAIL] Using helper functions in tests Gomega when assertions are used in the test body [It] prints USEFUL line number but MISSES automatic text in colorized text and summaries
  /ginkgo-test-helper/main_test.go:33
  [FAIL] Using helper functions in tests Gomega when the helper function returns an GomegaMatcher [It] prints USEFUL line number and USEFUL automatic text in colorized text and summaries. Getting the best of both worlds
  /ginkgo-test-helper/main_test.go:39

Ran 5 of 5 Specs in 0.002 seconds
FAIL! -- 0 Passed | 5 Failed | 0 Pending | 0 Skipped
--- FAIL: TestHelpers (0.00s)
FAIL
exit status 1
FAIL	ginkgo-test-helpers	0.159s
```
