package reflection

import (
  "reflect"
  "testing"
)

type Person struct {
  Name string
  Profile Profile
}

type Profile struct {
  Age int
  City string
}

func TestWalk(t *testing.T) {

  cases := []struct{
    Name string
    Input interface{}
    ExpectedCalls []string
  }{
    {
      "struct with one string field",
      struct{
        Name string
      }{"Wesley"},
      []string{"Wesley"},
    },
    {
      "struct with two string field",
      struct{
        Name string
        City string
      }{"Wesley", "Taipei"},
      []string{"Wesley", "Taipei"},
    },
    {
      "struct with non string field",
      struct{
        Name string
        Age int
      }{"Wesley", 24},
      []string{"Wesley"},
    },
    {
      "nested struct",
      Person{
        "Wesley",
        Profile{25, "Taipei"},
      },
      []string{"Wesley", "Taipei"},
    },
    {
      "pointer in nested struct",
      &Person{
        "Wesley",
        Profile{25, "Taipei"},
      },
      []string{"Wesley", "Taipei"},
    },
    {
      "pointer in nested struct",
      []Profile {
        {25, "Taipei"},
        {24, "Chunghua"},
      },
      []string{"Taipei", "Chunghua"},
    },
    {
      "arrays",
      [2]Profile {
        {25, "Taipei"},
        {24, "Chunghua"},
      },
      []string{"Taipei", "Chunghua"},
    },
    {
      "maps",
      map[string]string{
        "foo": "bar",
        "baz": "boz",
      },
      []string{"bar", "boz"},
    },
  }

  for _, test := range cases {
    t.Run(test.Name, func(t *testing.T) {
      var got []string
      walk(test.Input, func(s string) {
        got = append(got, s)
      })

      if !reflect.DeepEqual(got, test.ExpectedCalls){
        t.Errorf("got %v but want %v", got, test.ExpectedCalls)
      }
    })
  }

  t.Run("maps", func(t *testing.T) {
    aMap := map[string]string{
      "foo": "bar",
      "baz": "boz",
    }
    var got []string
    walk(aMap, func(input string) {
      got = append(got, input)
    })

    assertContains(t, got, "bar")
    assertContains(t, got, "boz")

  })

  t.Run("channel", func(t *testing.T) {
    aChannel := make(chan Profile)

    go func() {
      aChannel <- Profile{25, "Taipei"}
      aChannel <- Profile{24, "Chunghua"}
      close(aChannel)
    }()

    var got []string
    want := []string{"Taipei", "Chunghua"}

    walk(aChannel, func(input string) {
      got = append(got, input)
    })

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v but want %v", got, want)
    }
  })

  t.Run("function", func(t *testing.T) {
    aFunction := func() (Profile, Profile) {
      return Profile{25, "Taipei"}, Profile{24, "Chunghua"}
    }

    var got []string
    want := []string{"Taipei", "Chunghua"}

    walk(aFunction, func(input string) {
      got = append(got, input)
    })

    if !reflect.DeepEqual(got, want) {
      t.Errorf("got %v but want %v", got, want)
    }
  })
}

func assertContains(t testing.TB, got []string, value string) {
  t.Helper()
  contain := false
  for _, x := range got {
    if x == value {
      contain = true
    }
  }
  if !contain {
    t.Errorf("expected %v to contain %v but didn't", got, value)
  }

}
