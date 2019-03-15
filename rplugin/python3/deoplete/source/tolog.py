from deoplete.source.base import Base
class Source(Base):
    def __init__(self, vim):
        super().__init__(vim)
        self.name = 'tolog'
        self.mark = '[tolog]'
        self.rank = 1000
        self._count = 0
        self.min_pattern_length = 1
        self.matcher_key = "hoe"

    def gather_candidates(self, context):
        tags = self.vim.call("Tolog_Complete_tag")
        ret = tags.split("\n")
        print(ret)
        print("nyanko")
        ret.append("nyanko")

        return ret
