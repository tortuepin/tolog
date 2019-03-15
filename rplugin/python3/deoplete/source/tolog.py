from deoplete.source.base import Base
class Source(Base):
    def __init__(self, vim):
        super().__init__(vim)
        self.name = 'tolog'
        self.mark = '[tolog]'
        self.rank = 1000
        self._count = 0
        self.min_pattern_length = 1
        self.input_pattern = "@"
        self.filetype = ['markdown']

    def gather_candidates(self, context):
        tags = self.vim.call("Tolog_Complete_tag")
        ret = map(lambda x:x.strip("@"), tags.split("\n"))
        print(ret)

        return ret
