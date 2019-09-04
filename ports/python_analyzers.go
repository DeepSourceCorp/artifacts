/*
Language analyzers run multiple static check analyzers under the hood. In order
to maintain consistency, issues in one analyzer is suppressed in favour of a
similar issue in another analyzer. For example, F401 is an issue that is raised
by Flake8, but suppressed there and replaced by the same issue in Pylint, having
code W0614. In this case, if a user tries suppressing all issues in a line by
in the following way:

```python
from module import function  # noqa
# code that nowhere uses the imported function
```
or
```python
from module import function  # noqa: F401
# code that nowhere uses the imported function
```
the user would expect to not see the issue "Module imported but unused".

However, the message will be displayed, because in DeepSource platform, this issue
is raised using Pylint, for which the issue was not suppressed. In order to abstract
away this from the users, we need to support the analyzer specific issue silencers,
like `noqa` and `nosec`.
*/

package ports

/*
Issues that have been ported from other analyzers to Pylint
*/
type IssueSilencer struct {
	PortName     string            `json:"port_name"`
	SilencerCode string            `json:"silencer_code"`
	Issues       map[string]string `json:"issues"`
}

type LanguageMeta struct {
	CommentIdentifier string          `json:"comment_identifier"`
	Silencers         []IssueSilencer `json:"issue_silencers"`
}

var LanguagesMeta map[string]LanguageMeta

func init() {
	LanguagesMeta = map[string]LanguageMeta{
		".py": LanguageMeta{
			CommentIdentifier: "#",

			Silencers: []IssueSilencer{
				{
					PortName:     "Flake8ToPylint",
					SilencerCode: "noqa",
					Issues: map[string]string{
						"W0614": "F401",
					},
				},
				{
					PortName:     "BanditToPylint",
					SilencerCode: "nosec",
					Issues:       map[string]string{},
				},
			},
		},
		".go": LanguageMeta{
			CommentIdentifier: "//",

			Silencers: []IssueSilencer{},
		},
	}
}
